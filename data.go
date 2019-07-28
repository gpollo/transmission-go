package main

import "strconv"

type Method int

const (
	TorrentGet Method = iota
)

func (c Method) String() string {
	return [...]string{
		"torrent-get",
	}[c]
}

type Request struct {
	Method    string      `json:"method"`
	Arguments interface{} `json:"arguments"`
}

type Response struct {
	Arguments interface{} `json:"arguments"`
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

type Torrent struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Peers []Peer `json:"peers,omitempty"`
}

func (t *Torrent) fieldToString(field string) string {
	if field == "id" {
		return strconv.Itoa(t.ID)
	} else if field == "name" {
		return string(t.Name)
	} else if field == "peers" {
		return ""
	}

	return ""
}

type TorrentGetRequest struct {
	IDs    []int    `json:"ids,omitempty"`
	Fields []string `json:"fields"`
}

type TorrentGetResponse struct {
	Arguments struct {
		Torrents []Torrent `json:"torrents"`
	} `json:"arguments"`
}
