package packets

import (
	"fmt"
	"log"

	"github.com/m-nny/goinit/pkg/datatypes"
)

const (
	PACKET_ID_HANDSHAKE PacketID = 0x00
)

type Handshake struct {
	ProtocolVersion datatypes.VarInt
	ServerAddress   datatypes.String
	ServerPort      datatypes.UShort
	NextState       datatypes.VarInt
}

func NewHandshake(p *Packet) (*Handshake, error) {
	h := &Handshake{}
	log.Printf("packet: %+v", p)
	if err := p.Scan(&h.ProtocolVersion, &h.ServerAddress, &h.ServerPort, &h.NextState); err != nil {
		return nil, fmt.Errorf("could not scan handshake packet: %w", err)
	}
	if h.ProtocolVersion != PROTOCOL_VERSION {
		return h, fmt.Errorf("protocol version does not match server %v, client %v", PROTOCOL_VERSION, h.ProtocolVersion)
	}
	return h, nil
}
