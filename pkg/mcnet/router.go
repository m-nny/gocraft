package mcnet

import (
	"fmt"
	"io"
	"log"
)

type Router struct {
	handlers map[State]map[VarInt]Handler
}

func NewRouter() *Router {
	return &Router{
		handlers: make(map[State]map[VarInt]Handler),
	}
}

func (r *Router) AddRoute(state State, packetID VarInt, handler Handler) error {
	handlers := r.handlers[state]
	if handlers == nil {
		handlers = make(map[VarInt]Handler)
	}
	if _, ok := handlers[packetID]; ok {
		return fmt.Errorf("handler for state: %+v packetdID: %+v is already registered", state, packetID)
	}
	handlers[packetID] = handler
	r.handlers[state] = handlers
	return nil
}

func (router *Router) Handle(w ResponseWriter, r io.Reader, currentState State) error {
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
	PackgetLen   VarInt
	PacketID     VarInt
	CurrentState State
	Payload      []byte
}

type ResponseWriter interface {
	WriteJson(packetID VarInt, payload any) error
	WriteBytes(packetID VarInt, payload []byte) error
	SetState(newState State) error
}
