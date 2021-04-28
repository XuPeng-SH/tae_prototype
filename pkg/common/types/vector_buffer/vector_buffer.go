package vector_buffer

import (
	"bytes"
	"tae/pkg/common/types"
)

func New(options ...Option) *VectorBuffer {
	vf := &VectorBuffer{}
	for _, option := range options {
		*vf = option(*vf)
	}
	return vf
}

type Option func(VectorBuffer) VectorBuffer

func WithBufferType(vbt VectorBufferType) Option {
	return func(vf VectorBuffer) VectorBuffer {
		vf.Type = vbt
		return vf
	}
}

func WithSize(size int) Option {
	return func(vf VectorBuffer) VectorBuffer {
		vf.Buff = new(bytes.Buffer)
		vf.Buff.Grow(size)
		vf.Type = STANDARD_BUFFER
		return vf
	}
}

func WithItemType(itype types.PhysicalType) Option {
	return func(vf VectorBuffer) VectorBuffer {
		vf.ItemType = itype
		return vf
	}
}
