package datatypes

import "io"

type FieldDecoder io.ReaderFrom
type FieldEncoder io.WriterTo

type Field struct {
	FieldEncoder
	FieldDecoder
}
