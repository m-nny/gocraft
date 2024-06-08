package mcnet

import (
	"fmt"
	"io"
)

const PROTOCOL_VERSION = 766 // minecraft 1.20.6

type Packet interface {
	io.ReaderFrom
	PacketID() VarInt
}

const (
	PACKET_ID_HANDSHAKE VarInt = 0x00
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

func (p *HandshakePacket) PacketID() VarInt {
	return PACKET_ID_HANDSHAKE
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
