package main

import "fmt"
import "path"
import "os"

import "github.com/akamensky/argparse"

func main() {
	parser := argparse.NewParser(
		path.Base(os.Args[0]),
		"Command line interface for transmission RPC",
	)

	host := parser.String(
		"",
		"host",
		&argparse.Options{
			Required: false,
			Help:     "Hostname of the transmission RPC daemon",
		},
	)

	/* `list-torrents` command */

	cmdListTorrents := parser.NewCommand(
		"list-torrents",
		"Show the list of torrents",
	)

	cmdListTorrentsFields := cmdListTorrents.List(
		"f",
		"field",
		&argparse.Options{
			Required: false,
			Help:     "The list of fields to show",
		},
	)

	/* `list-files` command */

	cmdListFiles := parser.NewCommand(
		"list-files",
		"Show the list of files owned by a torrent",
	)

	cmdListFilesID := cmdListFiles.Int(
		"i",
		"id",
		&argparse.Options{
			Required: true,
			Help:     "The ID of the torrent",
		},
	)

	cmdListFilesFields := cmdListFiles.List(
		"f",
		"field",
		&argparse.Options{
			Required: false,
			Help:     "The list of fields to show",
		},
	)

	/* `rename-file` command */

	cmdRenameFile := parser.NewCommand(
		"rename-file",
		"Rename a file owned by a torrent",
	)

	cmdRenameFileID := cmdRenameFile.Int(
		"i",
		"id",
		&argparse.Options{
			Required: true,
			Help:     "The ID of the torrent",
		},
	)

	cmdRenameFileSrc := cmdRenameFile.String(
		"s",
		"src",
		&argparse.Options{
			Required: true,
			Help:     "The source file",
		},
	)

	cmdRenameFileDest := cmdRenameFile.String(
		"d",
		"dest",
		&argparse.Options{
			Required: true,
			Help:     "The destination file",
		},
	)

	/* arguments parsing */

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *host == "" {

	}

	client := Client{
		Endpoint: "http://192.168.0.108:9091/transmission/rpc",
	}

	if cmdListTorrents.Happened() {
		err := client.ListTorrents(*cmdListTorrentsFields)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if cmdListFiles.Happened() {
		err := client.ListFiles(*cmdListFilesID, *cmdListFilesFields)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if cmdRenameFile.Happened() {
		err := client.RenameFile(
			*cmdRenameFileID,
			*cmdRenameFileSrc,
			*cmdRenameFileDest,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	fmt.Println(*host)

	return
}
