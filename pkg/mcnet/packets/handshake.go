package packets

import (
	"fmt"
	"io"

	"github.com/m-nny/goinit/pkg/mcnet/datatypes"
)

var _ Packet = (*HandshakePacket)(nil)

const (
	PACKET_ID_HANDSHAKE PacketID = 0x00
)

type HandshakePacket struct {
	ProtocolVersion datatypes.VarInt
	ServerAddress   datatypes.String
	ServerPort      datatypes.UShort
	NextState       datatypes.VarInt
}

func (p *HandshakePacket) ReadFrom(r io.Reader) (int64, error) {
	nn, err := p.ProtocolVersion.ReadFrom(r)
	if err != nil {
		return nn, err
	}
	if p.ProtocolVersion != PROTOCOL_VERSION {
		return nn, fmt.Errorf("provided procotol version is not supported: want %d got %d", PROTOCOL_VERSION, p.ProtocolVersion)
	}
	if _, err := p.ServerAddress.ReadFrom(r); err != nil {
		return nn, err
	}
	if _, err := p.ServerPort.ReadFrom(r); err != nil {
		return nn, err
	}
	if _, err := p.NextState.ReadFrom(r); err != nil {
		return nn, err
	}
	return nn, nil
}
