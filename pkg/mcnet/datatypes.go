package mcnet

import (
	"fmt"
	"io"
)

type VarInt = int32

const VARINT_SEGMENT_BITS = 0x7F
const VARINT_CONTINUE_BIT = 0x80

func ParseVarInt(r io.Reader) (VarInt, error) {
	value := VarInt(0)
	pos := 0

	for {
		b := make([]byte, 1)
		if _, err := r.Read(b); err != nil {
			return value, err
		}
		value |= VarInt(b[0]&VARINT_SEGMENT_BITS) << pos
		if b[0]&VARINT_CONTINUE_BIT == 0 {
			break
		}
		pos += 7
		if pos >= 32 {
			return value, fmt.Errorf("VarInt too big")
		}
	}
	return value, nil
}

const STRING_MAX_LEN = 32767

func ParseString(r io.Reader) (string, error) {
	n, err := ParseVarInt(r)
	if err != nil {
		return "", err
	}
	if !(1 <= n && n <= STRING_MAX_LEN) {
		return "", fmt.Errorf("incorrect string length: want %d max: %d", n, STRING_MAX_LEN)
	}

	s := make([]byte, n)
	if _, err := io.ReadFull(r, s); err != nil {
		return "", err
	}
	return string(s), nil
}

type UShort = int16

func ParseUShort(r io.Reader) (UShort, error) {
	b := make([]byte, 2)
	if _, err := io.ReadFull(r, b); err != nil {
		return 0, err
	}
	value := UShort(b[0])<<8 | UShort(b[1])
	return value, nil
}
