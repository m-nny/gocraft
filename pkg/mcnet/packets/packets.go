package packets

import (
	"io"

	"github.com/m-nny/goinit/pkg/mcnet/datatypes"
)

const MINECRAFT_VERSION = "1.20.6"
const PROTOCOL_VERSION = 766 // minecraft 1.20.6

type PacketID = datatypes.VarInt

type Packet interface {
	io.ReaderFrom
}
