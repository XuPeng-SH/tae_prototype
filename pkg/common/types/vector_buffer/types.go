package vector_buffer

import (
	"tae/pkg/common/types"
)

type VectorBufferType uint8

const (
	STANDARD_BUFFER VectorBufferType = iota
	DICTIONARY_BUFFER
	VECTOR_CHILD_BUFFER
	STRING_BUFFER
	STRUCT_BUFFER
	LIST_BUFFER
)

type VectorBuffer struct {
	Type     VectorBufferType
	Data     []byte
	ItemType types.LogicType
	ItemSize uint8
}
