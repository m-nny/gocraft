package mcserver

import (
	"log"
	"net"

	"github.com/m-nny/goinit/pkg/datatypes"
	"github.com/m-nny/goinit/pkg/mcnet"
	"github.com/m-nny/goinit/pkg/packets"
)

var _ packets.ConnState = (*conn)(nil)

type conn struct {
	rwc    net.Conn
	state  datatypes.State
	router *mcnet.Router
}

func (c *conn) Close() error {
	return c.rwc.Close()
}

func (c *conn) SetState(newState datatypes.State) error {
	c.state = newState
	return nil
}

func (c *conn) GetState() datatypes.State {
	return c.state
}

func (c *conn) Serve() {
	for {
		if err := c.router.Handle(c, c.rwc); err != nil {
			log.Printf("[conn.serve] err: %v", err)
			c.Close()
			return
		}
	}
}
