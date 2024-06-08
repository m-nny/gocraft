package mcnet

import (
	"fmt"
	"io"
	"log"
)

const PROTOCOL_VERSION = 766 // minecraft 1.20.6

type Packet interface {
	io.ReaderFrom
	PacketID() VarInt
}

const (
	PACKET_ID_HANDSHAKE VarInt = 0x00
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

func ReadGenericPacket(r io.Reader) (Packet, error) {
	var packetLen VarInt
	if _, err := packetLen.ReadFrom(r); err != nil {
		return nil, err
	}
	log.Printf("got package with length: %d", packetLen)

	var packetID VarInt
	if _, err := packetID.ReadFrom(r); err != nil {
		return nil, err
	}
	log.Printf("got package with packetId: %d", packetID)

	var packet Packet

	if packetID == PACKET_ID_HANDSHAKE {
		packet = &HandshakePacket{}
	} else {
		return nil, fmt.Errorf("ReadPacket: packetId %d not implemented", packetID)
	}
	if _, err := packet.ReadFrom(r); err != nil {
		return nil, err
	}
	return packet, nil
}
