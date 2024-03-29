package main

import "fmt"
import "path"
import "os"

import "github.com/akamensky/argparse"

const (
	DefaultRPCEndpoint string = "http://host:9091/transmission/rpc"
)

func main() {
	parser := argparse.NewParser(
		path.Base(os.Args[0]),
		"Command line interface for transmission RPC",
	)

	argEndpoint := parser.String(
		"e",
		"endpoint",
		&argparse.Options{
			Required: false,
			Help:     "The endpoint of the transmission session",
		},
	)

	argNoHeader := parser.Flag(
		"",
		"no-header",
		&argparse.Options{
			Required: false,
			Help:     "Don't print table header",
			Default:  false,
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

	/* `rename-torrent` command */

	cmdRenameTorrent := parser.NewCommand(
		"rename-torrent",
		"Rename a torrent",
	)

	cmdRenameTorrentID := cmdRenameTorrent.Int(
		"i",
		"id",
		&argparse.Options{
			Required: true,
			Help:     "The ID of the torrent",
		},
	)

	cmdRenameTorrentName := cmdRenameTorrent.String(
		"n",
		"name",
		&argparse.Options{
			Required: true,
			Help:     "The new name of the torrent",
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

	commands := CreateCommands()
	commands.Client.Endpoint = DefaultRPCEndpoint
	commands.NoHeader = *argNoHeader

	envEndpoint := os.Getenv("TRANSMISSION_ENDPOINT")
	if envEndpoint != "" {
		commands.Client.Endpoint = envEndpoint
	}

	if *argEndpoint != "" {
		commands.Client.Endpoint = *argEndpoint
	}

	if cmdListTorrents.Happened() {
		err := commands.ListTorrents(*cmdListTorrentsFields)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if cmdListFiles.Happened() {
		err := commands.ListFiles(*cmdListFilesID, *cmdListFilesFields)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if cmdRenameTorrent.Happened() {
		err := commands.RenameTorrent(
			*cmdRenameTorrentID,
			*cmdRenameTorrentName,
		)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if cmdRenameFile.Happened() {
		err := commands.RenameFile(
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

	return
}
