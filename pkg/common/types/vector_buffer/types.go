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

type IVectorBuffer interface {
	Size() int
	MaxItems() int
	ReferenceOther(other interface{}, offset int)
	SetValue(idx int, val interface{})
	GetValue(idx int) interface{}
	GetItemType() types.LogicType
	GetType() VectorBufferType
}

type VectorBuffer struct {
	IVectorBuffer
	Type     VectorBufferType
	Data     []byte
	ItemType types.LogicType
	ItemSize uint8
}
