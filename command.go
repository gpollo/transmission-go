package main

import "net/http"
import "encoding/json"
import "github.com/olekukonko/tablewriter"
import "io/ioutil"
import "bytes"
import "errors"
import "os"
import "regexp"

const (
	SessionIDRegex string = "<code>X-Transmission-Session-Id: *(.*)</code>"
)

type Client struct {
	Endpoint  string
	SessionID string
}

func (c *Client) postRequest(payload []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", c.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Transmission-Session-Id", c.SessionID)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return []byte{}, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []byte{}, err
	}

	if response.StatusCode == 409 {
		regex := regexp.MustCompile(SessionIDRegex)

		matches := regex.FindStringSubmatch(string(body))
		if len(matches) == 2 {
			c.SessionID = matches[1]
		}

		return []byte{}, errors.New("Invalid Session ID")
	}

	if response.StatusCode != 200 {
		return []byte{}, errors.New(string(body))
	}

	return body, nil
}

func (c *Client) sendRequest(req interface{}) ([]byte, error) {
	parameters, err := json.Marshal(req)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := c.postRequest(parameters)
	if err != nil {
		bytes, err = c.postRequest(parameters)
		if err != nil {
			return []byte{}, err
		}
	}

	return bytes, nil
}

func (c *Client) ListTorrents(fields []string) error {
	fields = setAsFirstString(fields, "name")
	fields = setAsFirstString(fields, "id")

	request := TorrentGetRequest{}
	request.Method = TorrentGet.String()
	request.Arguments.Fields = fields

	bytes, err := c.sendRequest(request)
	if err != nil {
		return err
	}

	response := TorrentGetResponse{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fields)
	table.SetBorders(tablewriter.Border{
		Left:   false,
		Top:    false,
		Right:  false,
		Bottom: false,
	})
	table.SetCenterSeparator("-")
	table.SetColWidth(100)

	for _, torrent := range response.Arguments.Torrents {
		row := []string{}
		for _, field := range fields {
			row = append(row, torrent.fieldToString(field))
		}
		table.Append(row)
	}
	table.Render()

	return nil
}

func (c *Client) ListFiles(id int, fields []string) error {
	fields = setAsFirstString(fields, "name")

	request := TorrentGetRequest{}
	request.Method = TorrentGet.String()
	request.Arguments.IDs = []int{id}
	request.Arguments.Fields = []string{"files"}

	bytes, err := c.sendRequest(request)
	if err != nil {
		return err
	}

	response := TorrentGetResponse{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(fields)
	table.SetBorders(tablewriter.Border{
		Left:   true,
		Top:    true,
		Right:  false,
		Bottom: false,
	})
	table.SetCenterSeparator("-")
	table.SetColWidth(200)

	torrent := response.Arguments.Torrents[0]
	for _, file := range torrent.Files {
		row := []string{}
		for _, field := range fields {
			row = append(row, file.fieldToString(field))
		}
		table.Append(row)
	}
	table.Render()

	return nil
}
