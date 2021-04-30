package vector

import (
	"tae/pkg/common/types"
	"tae/pkg/common/types/value"
	// vmask "tae/pkg/common/types/validity_mask"
	vbuff "tae/pkg/common/types/vector_buffer"
)

func NewVector(options ...Option) *Vector {
	v := &Vector{
		Buff: vbuff.NewVectorBuffer(vbuff.WithItemType(types.LT_INVALID)),
		Type: FLAT_VECTOR,
	}
	for _, option := range options {
		*v = option(*v)
	}
	return v
}

type Option func(Vector) Vector

func WithInitByValue(val value.Value) Option {
	return func(vec Vector) Vector {
		// vec.Buff = vbuff.NewVectorBuffer(vbuff.WithSize())
		return vec
	}
}

func (vec *Vector) SetValue(idx int, val value.Value) {
	// TODO
}
