package main

import "github.com/olekukonko/tablewriter"
import "os"
import "strings"

type Commands struct {
	Client   Client
	NoHeader bool
}

func CreateCommands() Commands {
	commands := Commands{}
	commands.Client = CreateClient()
	commands.NoHeader = false

	return commands
}

func (c *Commands) printTable(header []string, content [][]string) {
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

func (c *Commands) ListTorrents(fields []string) error {
	fields = setAsFirstString(fields, "name")
	fields = setAsFirstString(fields, "id")

	response, err := c.Client.TorrentGet([]int{}, fields)
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

func (c *Commands) ListFiles(id int, fields []string) error {
	fields = setAsFirstString(fields, "name")

	response, err := c.Client.TorrentGet([]int{}, []string{"files"})
	if err != nil {
		return err
	}

	name, err := c.Client.GetTorrentStringParam(id, "name")
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

func (c *Commands) ListPeers(id int, fields []string) error {
	fields = setAsFirstString(fields, "address")

	response, err := c.Client.TorrentGet([]int{id}, []string{"peers"})
	if err != nil {
		return err
	}

	content := [][]string{}
	torrent := response.Arguments.Torrents[0]
	for _, peer := range torrent.Peers {
		row := []string{}
		for _, field := range fields {
			row = append(row, peer.fieldToString(field))
		}
		content = append(content, row)
	}
	c.printTable(fields, content)

	return nil
}

func (c *Commands) RenameTorrent(id int, new string) error {
	old, err := c.Client.GetTorrentStringParam(id, "name")
	if err != nil {
		return err
	}

	if _, err := c.Client.TorrentRenamePath(id, old, new); err != nil {
		return err
	}

	return nil
}

func (c *Commands) RenameFile(id int, src string, dest string) error {
	name, err := c.Client.GetTorrentStringParam(id, "name")
	if err != nil {
		return err
	}

	if _, err := c.Client.TorrentRenamePath(id, name+"/"+src, dest); err != nil {
		return err
	}

	return nil
}
