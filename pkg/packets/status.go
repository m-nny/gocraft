package packets

import (
	"encoding/json"

	"github.com/m-nny/goinit/pkg/datatypes"
)

const (
	PACKET_ID_STATUS PacketID = 0x00
)

func NewStatusResponsePacket() (*Packet, error) {
	type Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	}
	type Text struct {
		Text string `json:"text"`
	}
	type Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
	}
	type Response struct {
		Version     Version `json:"version"`
		Players     Players `json:"players"`
		Description Text    `json:"description,omitempty"`
	}
	response := &Response{
		Version: Version{
			Name:     "GoCraft",
			Protocol: PROTOCOL_VERSION,
		},
		Players: Players{
			Max:    1,
			Online: 0,
		},
		Description: Text{Text: "Custom server written in go"},
	}
	json, err := json.Marshal(response)
	if err != nil {
		return nil, err
	}
	payload := datatypes.String(json)

	return BuildPacket(PACKET_ID_STATUS, &payload)
}
