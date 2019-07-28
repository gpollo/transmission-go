package main

import "net/http"
import "fmt"
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

var sessionID string = ""

//var rpcURL string = ""

func postRequest(payload []byte) ([]byte, error) {
	request, err := http.NewRequest("POST", rpcURL, bytes.NewBuffer(payload))
	if err != nil {
		return []byte{}, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Transmission-Session-Id", sessionID)

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
			sessionID = matches[1]
		}

		return []byte{}, errors.New("Invalid Session ID")
	}

	return body, nil
}

func sendRequest(req interface{}) ([]byte, error) {
	parameters, err := json.Marshal(req)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := postRequest(parameters)
	if err != nil {
		bytes, err = postRequest(parameters)
		if err != nil {
			return []byte{}, err
		}
	}

	return bytes, nil
}

func ListTorrents(fields []string) error {
	request := TorrentGetRequest{}
	request.Method = TorrentGet.String()
	request.Arguments.Fields = fields

	bytes, err := sendRequest(request)
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

func ListFiles(id int) {

	fmt.Println(id)
}
