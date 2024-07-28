package packets_test

import (
	"bytes"
	"testing"

	"github.com/m-nny/goinit/pkg/packets"
	"github.com/stretchr/testify/require"
)

func Test_Packet_Unpack(t *testing.T) {
	testCases := []struct {
		name  string
		input []byte
		want  *packets.Packet
	}{
		{
			name:  "Handshake",
			input: []byte{16, 0, 254, 5, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 31, 144, 1},
			want: &packets.Packet{
				ID:   0,
				Data: []byte{254, 5, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 31, 144, 1},
			},
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
