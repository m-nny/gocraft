package mcnet

import (
	"fmt"
	"io"
	"log"
	"net"
)

type Client struct {
	conn  net.Conn
	state State
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn:  conn,
		state: STATE_HANDSHAKING,
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Welcome() error {
	if c.state != STATE_HANDSHAKING {
		return fmt.Errorf("client is not in handshaking state")
	}

	packet, err := ReadGenericPacket(c.conn, c.state)
	if err != nil {
		return fmt.Errorf("err reading packet: %v", err)
	}

	handshake, ok := packet.(*HandshakePacket)
	if !ok {
		return fmt.Errorf("expected Handshake Packet, got: %+v", packet)
	}
	log.Printf("got handshake packet: %+v", handshake)

	if handshake.NextState == 1 {
		c.state = STATE_STATUS
	} else if handshake.NextState == 2 {
		c.state = STATE_LOGIN
	} else if handshake.NextState == 3 {
		c.state = STATE_TRANSFER
	}

	packet, err = ReadGenericPacket(c.conn, c.state)
	if err != nil {
		return err
	}
	log.Printf("got packet: %+v", packet)

	return nil
}

func ReadGenericPacket(r io.Reader, state State) (Packet, error) {
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
	if state == STATE_HANDSHAKING {
		if packetID == PACKET_ID_HANDSHAKE {
			packet = &HandshakePacket{}
		}
	}
	if packet == nil {
		return nil, fmt.Errorf("ReadPacket: packetId %d with state %d not implemented", packetID, state)
	}
	if _, err := packet.ReadFrom(r); err != nil {
		return nil, err
	}
	return packet, nil
}
