package mcnet

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/m-nny/goinit/pkg/mcnet/datatypes"
)

var _ ResponseWriter = (*conn)(nil)

type conn struct {
	rwc    net.Conn
	state  datatypes.State
	router *Router
}

func (c *conn) Close() error {
	return c.rwc.Close()
}

func (c *conn) Welcome() error {
	return fmt.Errorf("not implemented")
}

func (c *conn) WriteJson(packetID datatypes.VarInt, payload any) error {
	bytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	return c.WriteBytes(packetID, bytes)
}

func (c *conn) WriteBytes(packetID datatypes.VarInt, payload []byte) error {
	return fmt.Errorf("not implemented")
}

func (c *conn) SetState(newState datatypes.State) error {
	c.state = newState
	return nil
}

func (c *conn) serve() {
	for {
		if err := c.router.Handle(c, c.rwc, c.state); err != nil {
			log.Printf("err: %v", err)
			c.Close()
			return
		}
	}
}
