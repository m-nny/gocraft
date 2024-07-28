package packets_test

// import (
// 	"bytes"
// 	"testing"
//
// 	"github.com/m-nny/goinit/pkg/packets"
// 	"github.com/stretchr/testify/require"
// )
//
// func Test_Handshake(t *testing.T) {
// 	testCases := []struct {
// 		name        string
// 		data        []byte
// 		wantPackage any
// 	}{
// 		{
// 			name: "HandshakePacket",
// 			data: []byte{254, 5, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 31, 144, 2},
// 			wantPackage: &packets.HandshakePacket{
// 				ProtocolVersion: 766,
// 				ServerAddress:   "localhost",
// 				ServerPort:      8080,
// 				NextState:       2,
// 			},
// 		},
// 	}
//
// 	for _, test := range testCases {
// 		t.Run(test.name, func(t *testing.T) {
// 			r := bytes.NewReader(test.data)
// 			got := &packets.HandshakePacket{}
// 			_, err := got.ReadFrom(r)
// 			require.NoError(t, err)
// 			require.Equal(t, test.wantPackage, got)
// 		})
// 	}
// }
//
// func Test_StatusResponse(t *testing.T) {
// 	testCases := []struct {
// 		name   string
// 		data   []byte
// 		packet *packets.StatusResponsePacket
// 	}{
// 		{
// 			name: "StatusResponse",
// 			data: []byte{0x2c, 0x7b, 0x22, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3a, 0x7b, 0x22, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x31, 0x2e, 0x32, 0x30, 0x2e, 0x36, 0x22, 0x2c, 0x22, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x3a, 0x37, 0x36, 0x36, 0x7d, 0x7d},
// 			packet: &packets.StatusResponsePacket{
// 				Version: struct {
// 					Name     string
// 					Protocol int
// 				}{
// 					Name:     "1.20.6",
// 					Protocol: 766,
// 				},
// 			},
// 		},
// 	}
// 	for _, test := range testCases {
// 		t.Run(test.name, func(t *testing.T) {
// 			w := &bytes.Buffer{}
// 			_, err := test.packet.WriteTo(w)
// 			require.NoError(t, err)
// 			require.Equal(t, test.data, w.Bytes())
// 		})
// 	}
// }
//
// // func Test_ReadGenericPacket(t *testing.T) {
// // 	testCases := []struct {
// // 		name        string
// // 		data        []byte
// // 		state       datatypes.State
// // 		wantPackage packets.Packet
// // 	}{
// // 		{
// // 			name:  "HandshakePacket",
// // 			data:  []byte{16, 0, 254, 5, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 31, 144, 2},
// // 			state: datatypes.STATE_HANDSHAKING,
// // 			wantPackage: &packets.HandshakePacket{
// // 				ProtocolVersion: 766,
// // 				ServerAddress:   "localhost",
// // 				ServerPort:      8080,
// // 				NextState:       2,
// // 			},
// // 		},
// // 		// {
// // 		// 	name:        "StatusPacket",
// // 		// 	data:        []byte{1, 0},
// // 		// 	state:       mcnet.STATE_STATUS,
// // 		// 	wantPackage: &mcnet.StatusRequestPacket{},
// // 		// },
// // 	}
// //
// // 	for _, test := range testCases {
// // 		t.Run(test.name, func(t *testing.T) {
// // 			r := bytes.NewReader(test.data)
// // 			gotPackage, gotErr := packets.ReadGenericPacket(r, test.state)
// // 			require.NoError(t, gotErr)
// // 			require.Equal(t, test.wantPackage, gotPackage)
// // 		})
// // 	}
// // }
