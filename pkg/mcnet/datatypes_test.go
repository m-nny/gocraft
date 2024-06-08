package mcnet_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/m-nny/goinit/pkg/mcnet"
	"github.com/stretchr/testify/require"
)

func Test_ParseVarInt(t *testing.T) {
	testCases := []struct {
		want  mcnet.VarInt
		bytes []byte
	}{
		{
			want:  0,
			bytes: []byte{0x00},
		},
		{
			want:  1,
			bytes: []byte{0x01},
		},
		{
			want:  2,
			bytes: []byte{0x02},
		},
		{
			want:  127,
			bytes: []byte{0x7f},
		},
		{
			want:  128,
			bytes: []byte{0x80, 0x01},
		},
		{
			want:  255,
			bytes: []byte{0xff, 0x01},
		},
		{
			want:  25565,
			bytes: []byte{0xdd, 0xc7, 0x01},
		},
		{
			want:  25565,
			bytes: []byte{0xdd, 0xc7, 0x01},
		},
		{
			want:  2097151,
			bytes: []byte{0xff, 0xff, 0x7f}},
		{
			want:  2147483647,
			bytes: []byte{0xff, 0xff, 0xff, 0xff, 0x07},
		},
		{
			want:  -1,
			bytes: []byte{0xff, 0xff, 0xff, 0xff, 0x0f},
		},
		{
			want:  -2147483648,
			bytes: []byte{0x80, 0x80, 0x80, 0x80, 0x08},
		},
	}
	for _, test := range testCases {
		name := fmt.Sprintf("%d", test.want)
		t.Run(name, func(t *testing.T) {
			r := bytes.NewReader(test.bytes)
			got, err := mcnet.ParseVarInt(r)
			require.NoError(t, err, "no error")
			require.Equal(t, test.want, got)
			require.Equal(t, r.Len(), 0, "no unread bytes")
		})
	}
}
