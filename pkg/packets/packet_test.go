package packets_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/m-nny/goinit/pkg/datatypes"
	"github.com/m-nny/goinit/pkg/packets"
	"github.com/stretchr/testify/require"
)

var (
	PACKET_HANDSHAKE_DATA = []byte{254, 5, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 31, 144, 1}
	PACKET_HANDSHAKE_RAW  = append([]byte{16, 0}, PACKET_HANDSHAKE_DATA...)
)

func Test_Packet_Unpack(t *testing.T) {
	testCases := []struct {
		name  string
		input []byte
		want  *packets.Packet
	}{
		{
			name:  "Handshake",
			input: PACKET_HANDSHAKE_RAW,
			want:  &packets.Packet{ID: 0, Data: PACKET_HANDSHAKE_DATA},
		},
		{
			name:  "Status request",
			input: []byte{1, 0},
			want: &packets.Packet{
				ID:   0,
				Data: []byte{},
			},
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := &packets.Packet{}
			err := got.Unpack(bytes.NewReader(test.input))
			require.NoError(t, err)
			require.Equal(t, test.want, got)
		})
	}
}

func Test_Packet_Scan(t *testing.T) {
	p := &packets.Packet{ID: 0, Data: PACKET_HANDSHAKE_DATA}
	var (
		ProtocolVersion datatypes.VarInt
		ServerAddress   datatypes.String
		ServerPort      datatypes.UShort
		NextState       datatypes.VarInt
	)
	err := p.Scan(&ProtocolVersion, &ServerAddress, &ServerPort, &NextState)
	require.NoError(t, err)
	fmt.Printf("%v %v %v %v\n", ProtocolVersion, ServerAddress, ServerPort, NextState)
}
