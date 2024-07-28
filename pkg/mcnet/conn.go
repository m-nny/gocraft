package mcnet

import (
	"encoding/json"
	"log"
	"net"

	"github.com/m-nny/goinit/pkg/mcnet/datatypes"
	mcnet "github.com/m-nny/goinit/pkg/mcnet/net"
)

var _ mcnet.ResponseWriter = (*conn)(nil)

type conn struct {
	rwc    net.Conn
	state  datatypes.State
	router *mcnet.Router
}

func (c *conn) Close() error {
	return c.rwc.Close()
}

func (c *conn) WriteJson(packetID datatypes.VarInt, payload any) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return c.WriteBytes(packetID, bytes)
}

func (c *conn) WriteBytes(packetID datatypes.VarInt, payload []byte) error {
	packetLength := datatypes.VarInt(datatypes.VarInt(packetID).Len() + len(payload))
	if _, err := packetLength.WriteTo(c.rwc); err != nil {
		return err
	}

	if _, err := packetID.WriteTo(c.rwc); err != nil {
		return err
	}

	if _, err := c.rwc.Write(payload); err != nil {
		return err
	}

	return nil
}

func (c *conn) SetState(newState datatypes.State) error {
	c.state = newState
	return nil
}

func (c *conn) serve() {
	for {
		if err := c.router.Handle(c, c.rwc, c.state); err != nil {
			log.Printf("[conn.serve] err: %v", err)
			c.Close()
			return
		}
	}
}
