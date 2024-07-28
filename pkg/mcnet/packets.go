package mcnet

import (
	"encoding/json"
	"fmt"
	"io"
)

const MINECRAFT_VERSION = "1.20.6"
const PROTOCOL_VERSION = 766 // minecraft 1.20.6

type Packet interface {
	io.ReaderFrom
}

const (
	PACKET_ID_HANDSHAKE VarInt = 0x00
	PACKET_ID_STATUS    VarInt = 0x00
)

type State int

const (
	STATE_UNKNOWN State = iota
	STATE_HANDSHAKING
	STATE_STATUS
	STATE_LOGIN
	STATE_TRANSFER
)

var _ Packet = (*HandshakePacket)(nil)

type HandshakePacket struct {
	ProtocolVersion VarInt
	ServerAddress   String
	ServerPort      UShort
	NextState       VarInt
}

func (p *HandshakePacket) ReadFrom(r io.Reader) (int64, error) {
	nn, err := p.ProtocolVersion.ReadFrom(r)
	if err != nil {
		return nn, err
	}
	if p.ProtocolVersion != PROTOCOL_VERSION {
		return nn, fmt.Errorf("provided procotol version is not supported: want %d got %d", PROTOCOL_VERSION, p.ProtocolVersion)
	}
	if _, err := p.ServerAddress.ReadFrom(r); err != nil {
		return nn, err
	}
	if _, err := p.ServerPort.ReadFrom(r); err != nil {
		return nn, err
	}
	if _, err := p.NextState.ReadFrom(r); err != nil {
		return nn, err
	}
	return nn, nil
}

var _ Packet = (*StatusRequestPacket)(nil)

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
	s := String(str)
	return s.WriteTo(w)
}
