package mcnet_test

import (
	"bytes"
	"testing"

	"github.com/m-nny/goinit/pkg/mcnet"
	"github.com/stretchr/testify/require"
)

func Test_ReadGenericPacket(t *testing.T) {
	testCases := []struct {
		name        string
		data        []byte
		wantPackage *mcnet.HandshakePacket
	}{
		{
			name: "HandshakePacket",
			data: []byte{16, 0, 254, 5, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 31, 144, 2},
			wantPackage: &mcnet.HandshakePacket{
				ProtocolVersion: 766,
				ServerAddress:   "localhost",
				ServerPort:      8080,
				NextState:       2,
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			r := bytes.NewReader(test.data)
			gotPackage, gotErr := mcnet.ReadGenericPacket(r)
			require.NoError(t, gotErr)
			require.Equal(t, test.wantPackage, gotPackage)
		})
	}
}
