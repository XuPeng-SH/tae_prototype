package vector

import (
	"tae/pkg/common/types"
	// vmask "tae/pkg/common/types/validity_mask"
	vbuff "tae/pkg/common/types/vector_buffer"
)

func NewVector(options ...Option) *Vector {
	v := &Vector{
		Buff: vbuff.NewVectorBuffer(vbuff.WithItemType(types.NA)),
		Type: FLAT_VECTOR,
	}
	for _, option := range options {
		*v = option(*v)
	}
	return v
}

type Option func(Vector) Vector
