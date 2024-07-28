package mcnet

import (
	"fmt"
	"io"
	"log"

	"github.com/m-nny/goinit/pkg/mcnet/datatypes"
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

func (router *Router) Handle(w ResponseWriter, r io.Reader, currentState datatypes.State) error {
	req := &Request{}
	if _, err := req.PackgetLen.ReadFrom(r); err != nil {
		return err
	}
	log.Printf("got package with length: %d", req.PackgetLen)

	if _, err := req.PacketID.ReadFrom(r); err != nil {
		return err
	}
	log.Printf("got package with packetId: %d", req.PacketID)

	handler := router.handlers[currentState][req.PacketID]
	if handler == nil {
		return fmt.Errorf("handler for req: %+v is not registered", req)
	}
	if err := handler(w, req); err != nil {
		return fmt.Errorf("error while handling req: %+v err: %w", req, err)
	}
	return nil
}

type Handler func(ResponseWriter, *Request) error

type Request struct {
	PackgetLen   datatypes.VarInt
	PacketID     datatypes.PacketID
	CurrentState datatypes.State
	Payload      []byte
}

type ResponseWriter interface {
	WriteJson(packetID datatypes.VarInt, payload any) error
	WriteBytes(packetID datatypes.VarInt, payload []byte) error
	SetState(newState datatypes.State) error
}
