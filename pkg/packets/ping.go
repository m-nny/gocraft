package packets

import (
	"fmt"
	"log"

	"github.com/m-nny/goinit/pkg/datatypes"
)

const (
	PACKET_ID_PING PacketID = 0x01
)

type Ping struct {
	Timestamp datatypes.Long
}

func NewPing(p *Packet) (*Ping, error) {
	h := &Ping{}
	log.Printf("packet: %+v", p)
	if err := p.Scan(&h.Timestamp); err != nil {
		return nil, fmt.Errorf("could not scan ping packet: %w", err)
	}
	return h, nil
}
