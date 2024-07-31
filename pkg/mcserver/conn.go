package mcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/m-nny/goinit/pkg/packets"
)

type State int

const (
	STATE_UNKNOWN State = iota
	STATE_STATUS
	STATE_LOGIN
	STATE_TRANSFER
)

type conn struct {
	rwc net.Conn
}

func (c *conn) Close() error {
	log.Printf("[conn.Close] stoping client %+v", c.rwc.RemoteAddr())
	return c.rwc.Close()
}

func (c *conn) Serve() {
	defer c.Close()
	p, err := packets.New(c.rwc)
	if err != nil {
		log.Printf("[conn.serve] err: %v", err)
		return
	}
	handshake, err := packets.NewHandshake(p)
	if err != nil {
		log.Printf("[conn.serve] err: %v", err)
		return
	}
	log.Printf("got handshake: %+v", handshake)
	switch nextStatus := State(handshake.NextState); nextStatus {
	case STATE_STATUS:
		if err := c.handleStatus(); err != nil {
			log.Printf("error handling status %d: %v", nextStatus, err)
			return
		}
	default:
		log.Printf("status %d not implemented", nextStatus)
	}
}

func (c *conn) handleStatus() error {
	for i := 0; i < 2; i++ {
		p, err := packets.New(c.rwc)
		if err != nil {
			return err
		}
		log.Printf("[conn.handleStatus]: got status packet #%d %+v", i+1, p)
		if p.ID == packets.PACKET_ID_STATUS {
			log.Printf("[conn.handleStatus: PACKET_ID_STATUS")
			response, err := packets.NewStatusResponsePacket()
			if err != nil {
				return err
			}
			if err := c.Respond(response); err != nil {
				return err
			}
			// } else if p.ID == packets.PACKET_ID_STATUS {
			// 	return fmt.Errorf("not implemented")
		} else {
			return fmt.Errorf("packetID %d not implemented", p.ID)
		}
	}
	return fmt.Errorf("not implemented")
}

func (c *conn) Respond(packet *packets.Packet) error {
	return packet.Pack(c.rwc)
}
