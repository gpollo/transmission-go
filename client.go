package main

import "net/http"
import "encoding/json"
import "io/ioutil"
import "bytes"
import "errors"
import "regexp"

const (
	SessionIDRegex string = "<code>X-Transmission-Session-Id: *(.*)</code>"
)

type Client struct {
	Endpoint  string
	SessionID string
}

func CreateClient() Client {
	client := Client{}

	return client
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

func (c *Client) TorrentGet(ids []int, fields []string) (TorrentGetResponse, error) {
	request := TorrentGetRequest{}
	request.Method = TorrentGet.String()
	request.Arguments.IDs = ids
	request.Arguments.Fields = fields

	bytes, err := c.sendRequest(request)
	if err != nil {
		return TorrentGetResponse{}, err
	}

	response := TorrentGetResponse{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		return TorrentGetResponse{}, err
	}

	return response, nil
}

func (c *Client) TorrentRenamePath(id int, old string, new string) error {
	request := TorrentRenamePathRequest{}
	request.Method = TorrentRenamePath.String()
	request.Arguments.ID = id
	request.Arguments.Path = old
	request.Arguments.Name = new

	bytes, err := c.sendRequest(request)
	if err != nil {
		return err
	}

	response := TorrentRenamePathResponse{}
	if err := json.Unmarshal(bytes, &response); err != nil {
		return err
	}

	return nil
}
