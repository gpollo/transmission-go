package main

import "strconv"

type Method int

const (
	TorrentGet Method = iota
	TorrentRenamePath
)

func (c Method) String() string {
	return [...]string{
		"torrent-get",
		"torrent-rename-path",
	}[c]
}

type Peer struct {
	Address            string  `json:"address,omitempty"`
	ClientName         string  `json:"clientName,omitempty"`
	ClientIsChoked     bool    `json:"clientIsChoked,omitempty"`
	ClientIsInterested bool    `json:"clientIsInterested,omitempty"`
	FlagString         string  `json:"flagStr,omitempty"`
	IsDownloadingFrom  bool    `json:"isDownloadingFrom,omitempty"`
	IsEncrypted        bool    `json:"isEncrypted,omitempty"`
	IsIncoming         bool    `json:"isIncoming,omitempty"`
	IsUploadingTo      bool    `json:"isUploadingTo,omitempty"`
	IsUTP              bool    `json:"isUTP,omitempty"`
	PeerIsChoked       bool    `json:"peerIsChoked,omitempty"`
	PeerIsInterested   bool    `json:"peerIsInterested,omitempty"`
	Port               int     `json:"port,omitempty"`
	Progress           float32 `json:"progress,omitempty"`
	RateToClient       int     `json:"rateToClient,omitempty"`
	RateToPeer         int     `json:"rateToPeer,omitempty"`
}

type File struct {
	BytesCompleted int    `json:"bytesCompleted,omitempty"`
	Length         int    `json:"length,omitempty"`
	Name           string `json:"name,omitempty"`
}

func (f *File) fieldToString(field string) string {
	switch field {
	case "bytesCompleted":
		return strconv.Itoa(f.BytesCompleted)
	case "length":
		return strconv.Itoa(f.Length)
	case "name":
		return string(f.Name)
	}

	return "..."
}

type Torrent struct {
	Files []File `json:"files,omitempty"`
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Peers []Peer `json:"peers,omitempty"`
}

func (t *Torrent) fieldToString(field string) string {
	switch field {
	case "files":
		return "..."
	case "id":
		return strconv.Itoa(t.ID)
	case "name":
		return t.Name
	case "peers":
		return "..."
	}

	return "..."
}

type TorrentGetRequest struct {
	Method    string `json:"method"`
	Arguments struct {
		IDs    []int    `json:"ids,omitempty"`
		Fields []string `json:"fields"`
	} `json:"arguments"`
}

type TorrentGetResponse struct {
	Arguments struct {
		Torrents []Torrent `json:"torrents"`
	} `json:"arguments"`
}

type TorrentRenamePathRequest struct {
	Method    string `json:"method"`
	Arguments struct {
		ID   int    `json:"ids"`
		Path string `json:"path"`
		Name string `json:"name"`
	} `json:"arguments"`
}

type TorrentRenamePathResponse struct {
	Method    string `json:"method"`
	Arguments struct {
		ID   int    `json:"id"`
		Path string `json:"path"`
		Name string `json:"name"`
	} `json:"arguments"`
}
