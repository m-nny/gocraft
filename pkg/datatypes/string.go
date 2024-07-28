package datatypes

import (
	"fmt"
	"io"
)

const STRING_MAX_LEN = 32767

type String string

func (value *String) ReadFrom(r io.Reader) (int64, error) {
	var strlen VarInt
	nn, err := strlen.ReadFrom(r)
	if err != nil {
		return nn, err
	}
	if !(1 <= strlen && strlen <= STRING_MAX_LEN) {
		return nn, fmt.Errorf("incorrect string length: want %d max: %d", strlen, STRING_MAX_LEN)
	}
	s := make([]byte, strlen)
	n, err := io.ReadFull(r, s)
	nn += int64(n)
	if err != nil {
		return nn, err
	}
	*value = String(s)
	return nn, nil
}

func (value *String) WriteTo(w io.Writer) (int64, error) {
	strlen := VarInt(len(*value))
	if !(1 <= strlen && strlen <= STRING_MAX_LEN) {
		return 0, fmt.Errorf("incorrect string length: want %d max: %d", strlen, STRING_MAX_LEN)
	}
	nn, err := strlen.WriteTo(w)
	if err != nil {
		return nn, err
	}

	n, err := w.Write([]byte(*value))
	nn += int64(n)
	if err != nil {
		return nn, err
	}
	return nn, nil
}
