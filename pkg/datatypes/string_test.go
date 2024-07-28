package datatypes_test

import (
	"bytes"
	"testing"

	"github.com/m-nny/goinit/pkg/datatypes"
	"github.com/stretchr/testify/require"
)

func Test_String(t *testing.T) {
	testCases := []datatypes.String{
		"localhost",
		"abc",
		"1",
	}
	for _, test := range testCases {
		t.Run(string(test), func(t *testing.T) {
			buf := &bytes.Buffer{}
			_, err := test.WriteTo(buf)
			require.NoError(t, err)

			var got datatypes.String

			_, err = got.ReadFrom(buf)
			require.NoError(t, err)

			require.Equal(t, test, got)

		})
	}
}
