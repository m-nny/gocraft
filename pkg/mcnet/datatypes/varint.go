package datatypes

import (
	"fmt"
	"io"
)

const VARINT_SEGMENT_BITS = 0x7F
const VARINT_CONTINUE_BIT = 0x80

type VarInt int32

func (value *VarInt) ReadFrom(r io.Reader) (int64, error) {
	pos := 0
	nn := int64(0)
	*value = 0
	b := make([]byte, 1)
	for {
		n, err := r.Read(b)
		nn += int64(n)
		if err != nil {
			return nn, err
		}
		*value |= VarInt(b[0]&VARINT_SEGMENT_BITS) << pos
		if b[0]&VARINT_CONTINUE_BIT == 0 {
			break
		}
		pos += 7
		if pos >= 32 {
			return nn, fmt.Errorf("VarInt too big")
		}
	}
	return nn, nil
}

func (v VarInt) WriteTo(w io.Writer) (int64, error) {
	nn := int64(0)
	value := uint32(v)
	for {
		b := byte(value) & VARINT_SEGMENT_BITS
		value >>= 7
		if value != 0 {
			b |= VARINT_CONTINUE_BIT
		}

		n, err := w.Write([]byte{b})
		nn += int64(n)
		if err != nil {
			return nn, err
		}
		if value == 0 {
			break
		}
	}
	return nn, nil
}
