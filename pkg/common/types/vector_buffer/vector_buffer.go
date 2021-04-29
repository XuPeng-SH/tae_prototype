package vector_buffer

import (
	"tae/pkg/common/types"
)

func New(options ...Option) *VectorBuffer {
	vf := &VectorBuffer{
		Type: STANDARD_BUFFER,
	}
	for _, option := range options {
		*vf = option(*vf)
	}
	return vf
}

type Option func(VectorBuffer) VectorBuffer

func WithSize(size int) Option {
	return func(vf VectorBuffer) VectorBuffer {
		if size < 0 {
			panic("")
		}
		vf.Data = make([]byte, 0, size)
		return vf
	}
}

func WithItemType(itype types.PhysicalType) Option {
	return func(vf VectorBuffer) VectorBuffer {
		vf.ItemType = itype
		return vf
	}
}

func (vf *VectorBuffer) GetType() VectorBufferType {
	return vf.Type
}
