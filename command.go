package main

import "github.com/olekukonko/tablewriter"
import "os"
import "strings"
import "errors"

func (c *Client) printTable(header []string, content [][]string) {
	table := tablewriter.NewWriter(os.Stdout)

	if !c.NoHeader {
		table.SetHeader(header)
	}

	table.SetBorders(tablewriter.Border{
		Left:   false,
		Top:    false,
		Right:  false,
		Bottom: false,
	})
	table.SetColumnSeparator("")
	table.SetCenterSeparator("-")
	table.SetColWidth(100)
	table.AppendBulk(content)
	table.Render()
}

func (c *Client) ListTorrents(fields []string) error {
	fields = setAsFirstString(fields, "name")
	fields = setAsFirstString(fields, "id")

	response, err := c.TorrentGet([]int{}, fields)
	if err != nil {
		return err
	}

	content := [][]string{}
	for _, torrent := range response.Arguments.Torrents {
		row := []string{}
		for _, field := range fields {
			row = append(row, torrent.fieldToString(field))
		}
		content = append(content, row)
	}
	c.printTable(fields, content)

	return nil
}

func (c *Client) GetTorrentParam(id int, param string) (string, error) {
	response, err := c.TorrentGet([]int{id}, []string{param})
	if err != nil {
		return "", err
	}

	if len(response.Arguments.Torrents) != 1 {
		return "", errors.New("Unexpected number of torrent received")
	}

	return response.Arguments.Torrents[0].fieldToString(param), nil
}

func (c *Client) ListFiles(id int, fields []string) error {
	fields = setAsFirstString(fields, "name")

	response, err := c.TorrentGet([]int{}, []string{"files"})
	if err != nil {
		return err
	}

	name, err := c.GetTorrentParam(id, "name")
	if err != nil {
		return err
	}
	name = name + "/"

	content := [][]string{}
	torrent := response.Arguments.Torrents[0]
	for _, file := range torrent.Files {
		row := []string{}
		for _, field := range fields {
			if field == "name" {
				filename := file.Name
				filename = strings.TrimPrefix(filename, name)
				row = append(row, filename)
			} else {
				row = append(row, file.fieldToString(field))
			}
		}
		content = append(content, row)
	}
	c.printTable(fields, content)

	return nil
}

func (c *Client) RenameTorrent(id int, new string) error {
	old, err := c.GetTorrentParam(id, "name")
	if err != nil {
		return err
	}

	if err := c.TorrentRenamePath(id, old, new); err != nil {
		return err
	}

	return nil
}

func (c *Client) RenameFile(id int, src string, dest string) error {
	name, err := c.GetTorrentParam(id, "name")
	if err != nil {
		return err
	}

	if err := c.TorrentRenamePath(id, name+"/"+src, dest); err != nil {
		return err
	}

	return nil
}
