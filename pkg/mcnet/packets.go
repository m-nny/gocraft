package mcnet

import (
	"fmt"
	"io"
	"log"
)

const PROTOCOL_VERSION = 766 // minecraft 1.20.6

const (
	PACKET_ID_HANDSHAKE VarInt = 0x00
)

type HandshakePacket struct {
	ProtocolVersion VarInt
	ServerAddress   string
	ServerPort      UShort
	NextState       VarInt
}

func readHandshakePacketData(r io.Reader) (*HandshakePacket, error) {
	protocolVersion, err := ParseVarInt(r)
	if err != nil {
		return nil, err
	}
	if protocolVersion != PROTOCOL_VERSION {
		return nil, fmt.Errorf("provided procotol version is not supported: want %d got %d", PROTOCOL_VERSION, protocolVersion)
	}

	serverAddress, err := ParseString(r)
	if err != nil {
		return nil, err
	}

	serverPort, err := ParseUShort(r)
	if err != nil {
		return nil, err
	}

	nextState, err := ParseVarInt(r)
	if err != nil {
		return nil, err
	}

	return &HandshakePacket{
		ProtocolVersion: protocolVersion,
		ServerAddress:   serverAddress,
		ServerPort:      serverPort,
		NextState:       nextState,
	}, nil
}

func ReadHandshakePacket(r io.Reader) (*HandshakePacket, error) {
	packetLen, err := ParseVarInt(r)
	if err != nil {
		return nil, err
	}
	log.Printf("got package with length: %d", packetLen)

	packetID, err := ParseVarInt(r)
	if err != nil {
		return nil, err
	}
	log.Printf("got package with packetId: %d", packetID)

	if packetID != PACKET_ID_HANDSHAKE {
		return nil, fmt.Errorf("expected handshake packet, got packetID: %d", packetID)
	}

	return readHandshakePacketData(r)
}

func ReadGenericPacket(r io.Reader) (any, error) {
	packetLen, err := ParseVarInt(r)
	if err != nil {
		return nil, err
	}
	log.Printf("got package with length: %d", packetLen)

	packetID, err := ParseVarInt(r)
	if err != nil {
		return nil, err
	}
	log.Printf("got package with packetId: %d", packetID)

	if packetID == PACKET_ID_HANDSHAKE {
		return readHandshakePacketData(r)
	}

	return nil, fmt.Errorf("ReadPacket: packetId %d not implemented", packetID)
}
