package main

import "strconv"

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

func (p *Peer) fieldToString(field string) string {
	return "..."
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
