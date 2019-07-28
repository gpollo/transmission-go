package main

import "fmt"
import "path"
import "os"

//import "encoding/json"
import "github.com/akamensky/argparse"

//import "net/http"

var rpcURL = "http://192.168.0.108:9091/transmission/rpc"

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

	if err := parser.Parse(os.Args); err != nil {
		fmt.Print(parser.Usage(err))
	}

	if *host == "" {

	}

	if cmdListTorrents.Happened() {
		err := ListTorrents(*cmdListTorrentsFields)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	if cmdListFiles.Happened() {
		err := ListFiles(*cmdListFilesID, *cmdListFilesFields)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	fmt.Println(*host)

	return
}
