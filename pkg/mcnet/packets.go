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
	ServerAddress   String
	ServerPort      UShort
	NextState       VarInt
}

func readHandshakePacketData(r io.Reader) (*HandshakePacket, error) {
	var protocolVersion VarInt
	if _, err := protocolVersion.ReadFrom(r); err != nil {
		return nil, err
	}
	if protocolVersion != PROTOCOL_VERSION {
		return nil, fmt.Errorf("provided procotol version is not supported: want %d got %d", PROTOCOL_VERSION, protocolVersion)
	}

	var serverAddress String
	if _, err := serverAddress.ReadFrom(r); err != nil {
		return nil, err
	}

	var serverPort UShort
	if _, err := serverPort.ParseUShort(r); err != nil {
		return nil, err
	}

	var nextState VarInt
	if _, err := nextState.ReadFrom(r); err != nil {
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

	if packetID != PACKET_ID_HANDSHAKE {
		return nil, fmt.Errorf("expected handshake packet, got packetID: %d", packetID)
	}

	return readHandshakePacketData(r)
}

func ReadGenericPacket(r io.Reader) (any, error) {
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

	if packetID == PACKET_ID_HANDSHAKE {
		return readHandshakePacketData(r)
	}

	return nil, fmt.Errorf("ReadPacket: packetId %d not implemented", packetID)
}
