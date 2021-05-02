package vbuff

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
	ToString(opts ...interface{}) string
	Size() types.IDX_T
	MaxItems() types.IDX_T
	ReferenceOther(other IVectorBuffer, offset types.IDX_T)
	SetValue(idx types.IDX_T, val interface{})
	GetValue(idx types.IDX_T) interface{}
	GetItemType() types.LogicType
	GetItemSize() types.IDX_T
	GetData() []byte
	GetType() VectorBufferType
	ForceRepeat(from_idx, count types.IDX_T)
}

type VectorBuffer struct {
	// IVectorBuffer
	Type     VectorBufferType
	Data     []byte
	ItemType types.LogicType
	ItemSize types.IDX_T
}
