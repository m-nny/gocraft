package packets

import (
	"io"

	"github.com/m-nny/goinit/pkg/datatypes"
)

const MINECRAFT_VERSION = "1.20.6"
const PROTOCOL_VERSION = 766 // minecraft 1.20.6

type PacketID = datatypes.VarInt

type Packet struct {
	ID   PacketID
	Data []byte
}

func (p *Packet) Unpack(r io.Reader) error {
	var packetLen datatypes.VarInt
	if _, err := packetLen.ReadFrom(r); err != nil {
		return err
	}
	var packetID datatypes.VarInt
	n, err := packetID.ReadFrom(r)
	if err != nil {

		return err
	}
	p.ID = packetID

	dataLen := int64(packetLen) - n
	// TODO: reuse potentially existing p.data
	p.Data = make([]byte, dataLen)
	if _, err := io.ReadFull(r, p.Data); err != nil {
		return err
	}
	return nil
}

func (p *Packet) Pack(w io.Writer) error {
	packetLength := datatypes.VarInt(p.ID.Len() + len(p.Data))
	if _, err := packetLength.WriteTo(w); err != nil {
		return err
	}

	if _, err := p.ID.WriteTo(w); err != nil {
		return err
	}

	if _, err := w.Write(p.Data); err != nil {
		return err
	}

	return nil
}

type ConnState interface {
	SetState(newState datatypes.State) error
	GetState() datatypes.State
}
