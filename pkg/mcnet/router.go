package mcnet

import (
	"fmt"
	"io"

	"github.com/m-nny/goinit/pkg/datatypes"
	"github.com/m-nny/goinit/pkg/packets"
)

type Router struct {
	handlers map[datatypes.State]map[datatypes.VarInt]Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[datatypes.State]map[datatypes.VarInt]Handler),
	}
}

func (r *Router) AddRoute(state datatypes.State, packetID datatypes.VarInt, handler Handler) error {
	handlers := r.handlers[state]
	if handlers == nil {
		handlers = make(map[datatypes.VarInt]Handler)
	}
	if _, ok := handlers[packetID]; ok {
		return fmt.Errorf("handler for state: %+v packetdID: %+v is already registered", state, packetID)
	}
	handlers[packetID] = handler
	r.handlers[state] = handlers
	return nil
}

func (router *Router) Handle(cs packets.ConnState, rw io.ReadWriter) error {
	packet := &packets.Packet{}
	if err := packet.Unpack(rw); err != nil {
		return err
	}

	handler := router.handlers[cs.GetState()][packet.ID]
	if handler == nil {
		return fmt.Errorf("handler for req: %+v %+v is not registered", cs.GetState(), packet.ID)
	}
	if err := handler(cs, packet); err != nil {
		return fmt.Errorf("error while handling req: %+v %+v err: %w", cs.GetState(), packet.ID, err)
	}
	return nil
}

type Handler func(packets.ConnState, *packets.Packet) error
