package datatypes

import (
	"io"
)

type Long int64

func (value *Long) ReadFrom(r io.Reader) (int64, error) {
	b := make([]byte, 8)
	n, err := io.ReadFull(r, b)
	nn := int64(n)
	if err != nil {
		return nn, err
	}
	*value = Long(b[0])<<24 | Long(b[1])<<16 | Long(b[2])<<8 | Long(b[3])
	return nn, nil
}
