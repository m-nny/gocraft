package datatypes

import (
	"io"
)

type UShort int16

func (value *UShort) ReadFrom(r io.Reader) (int64, error) {
	b := make([]byte, 2)
	n, err := io.ReadFull(r, b)
	nn := int64(n)
	if err != nil {
		return nn, err
	}
	*value = UShort(b[0])<<8 | UShort(b[1])
	return nn, nil
}
