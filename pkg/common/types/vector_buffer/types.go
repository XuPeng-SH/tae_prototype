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

func (vt VectorBufferType) String() string {
	switch vt {
	case STANDARD_BUFFER:
		return "STANDARD_BUFFER"
	case DICTIONARY_BUFFER:
		return "DICTIONARY_BUFFER"
	case VECTOR_CHILD_BUFFER:
		return "VECTOR_CHILD_BUFFER"
	case STRING_BUFFER:
		return "STRING_BUFFER"
	case STRUCT_BUFFER:
		return "STRUCT_BUFFER"
	case LIST_BUFFER:
		return "LIST_BUFFER"
	}
	panic("")
}

type IVectorBuffer interface {
	String() string
	Size() int
	MaxItems() int
	ReferenceOther(other IVectorBuffer, offset int)
	SetValue(idx int, val interface{})
	GetValue(idx int) interface{}
	GetItemType() types.LogicType
	GetItemSize() uint8
	GetData() []byte
	GetType() VectorBufferType
}

type VectorBuffer struct {
	IVectorBuffer
	Type     VectorBufferType
	Data     []byte
	ItemType types.LogicType
	ItemSize uint8
}
