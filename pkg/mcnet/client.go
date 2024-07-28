package mcnet

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/m-nny/goinit/pkg/mcnet/datatypes"
)

type Client struct {
	conn  net.Conn
	state datatypes.State
}

func NewClient(conn net.Conn) *Client {
	return &Client{
		conn:  conn,
		state: datatypes.STATE_HANDSHAKING,
	}
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Welcome() error {
	if c.state != datatypes.STATE_HANDSHAKING {
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
		c.state = datatypes.STATE_STATUS
	} else if handshake.NextState == 2 {
		c.state = datatypes.STATE_LOGIN
	} else if handshake.NextState == 3 {
		c.state = datatypes.STATE_TRANSFER
	}

	packet, err = ReadGenericPacket(c.conn, c.state)
	if err != nil {
		return err
	}
	log.Printf("got packet: %+v", packet)

	return nil
}

func ReadGenericPacket(r io.Reader, state datatypes.State) (Packet, error) {
	var packetLen datatypes.VarInt
	if _, err := packetLen.ReadFrom(r); err != nil {
		return nil, err
	}
	log.Printf("got package with length: %d", packetLen)

	var packetID datatypes.VarInt
	if _, err := packetID.ReadFrom(r); err != nil {
		return nil, err
	}
	log.Printf("got package with packetId: %d", packetID)

	var packet Packet
	if state == datatypes.STATE_HANDSHAKING {
		if packetID == datatypes.PACKET_ID_HANDSHAKE {
			packet = &HandshakePacket{}
		}
	} else if state == datatypes.STATE_STATUS {
		if packetID == datatypes.PACKET_ID_STATUS {
			packet = &StatusRequestPacket{}
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
