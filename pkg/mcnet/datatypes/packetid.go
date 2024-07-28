package datatypes

type PacketID = VarInt

const (
	PACKET_ID_HANDSHAKE PacketID = 0x00
	PACKET_ID_STATUS    PacketID = 0x00
)
