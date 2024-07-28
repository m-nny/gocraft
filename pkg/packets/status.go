package packets

import (
	"encoding/json"
	"io"

	"github.com/m-nny/goinit/pkg/datatypes"
)

const (
	PACKET_ID_STATUS PacketID = 0x00
)

type StatusRequestPacket struct{}

func (p *StatusRequestPacket) ReadFrom(r io.Reader) (int64, error) {
	return 0, nil
}

type StatusResponsePacket struct {
	Version struct {
		Name     string
		Protocol int
	}
	// Players struct {
	// 	Max    int `json:",omitempty"`
	// 	Online int `json:",omitempty"`
	// 	Sample []struct {
	// 		Name string
	// 		Id   string
	// 	} `json:",omitempty"`
	// } `json:",omitempty"`
	// Description struct {
	// 	Text string
	// } `json:",omitempty"`
	// FavIcon            string `json:",omitempty"`
	// EnforcesSecureChat bool   `json:",omitempty"`
	// PreviousChat       bool   `json:",omitempty"`
}

func NewStatusResponsePacket() *StatusResponsePacket {
	return &StatusResponsePacket{
		Version: struct {
			Name     string
			Protocol int
		}{
			Name:     MINECRAFT_VERSION,
			Protocol: PROTOCOL_VERSION,
		},
	}
}

func (p *StatusResponsePacket) WriteTo(w io.Writer) (int64, error) {
	str, err := json.Marshal(p)
	if err != nil {
		return 0, err
	}
	s := datatypes.String(str)
	return s.WriteTo(w)
}
