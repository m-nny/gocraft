package mcnet_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/m-nny/goinit/pkg/mcnet"
	"github.com/stretchr/testify/require"
)

var VarIntTestCases = []struct {
	value mcnet.VarInt
	bytes []byte
}{
	{
		value: 0,
		bytes: []byte{0x00},
	},
	{
		value: 1,
		bytes: []byte{0x01},
	},
	{
		value: 2,
		bytes: []byte{0x02},
	},
	{
		value: 127,
		bytes: []byte{0x7f},
	},
	{
		value: 128,
		bytes: []byte{0x80, 0x01},
	},
	{
		value: 255,
		bytes: []byte{0xff, 0x01},
	},
	{
		value: 25565,
		bytes: []byte{0xdd, 0xc7, 0x01},
	},
	{
		value: 25565,
		bytes: []byte{0xdd, 0xc7, 0x01},
	},
	{
		value: 2097151,
		bytes: []byte{0xff, 0xff, 0x7f}},
	{
		value: 2147483647,
		bytes: []byte{0xff, 0xff, 0xff, 0xff, 0x07},
	},
	{
		value: -1,
		bytes: []byte{0xff, 0xff, 0xff, 0xff, 0x0f},
	},
	{
		value: -2147483648,
		bytes: []byte{0x80, 0x80, 0x80, 0x80, 0x08},
	},
}

func Test_VarInt_ReadFrom(t *testing.T) {
	for _, test := range VarIntTestCases {
		name := fmt.Sprintf("%d", test.value)
		t.Run(name, func(t *testing.T) {
			r := bytes.NewReader(test.bytes)
			var got mcnet.VarInt
			_, err := got.ReadFrom(r)
			require.NoError(t, err, "no error")
			require.Equal(t, test.value, got)
			require.Equal(t, r.Len(), 0, "no unread bytes")
		})
	}
}

func Test_VarInt_WriteTo(t *testing.T) {
	for _, test := range VarIntTestCases {
		name := fmt.Sprintf("%d", test.value)
		t.Run(name, func(t *testing.T) {
			var w bytes.Buffer

			_, err := test.value.WriteTo(&w)
			got := w.Bytes()

			require.NoError(t, err, "no error")
			require.Equal(t, test.bytes, got)
		})
	}
}

func Test_String(t *testing.T) {
	testCases := []mcnet.String{
		"localhost",
		"abc",
		"1",
	}
	for _, test := range testCases {
		t.Run(string(test), func(t *testing.T) {
			buf := &bytes.Buffer{}
			_, err := test.WriteTo(buf)
			require.NoError(t, err)

			var got mcnet.String

			_, err = got.ReadFrom(buf)
			require.NoError(t, err)

			require.Equal(t, test, got)

		})
	}
}
