package selection_vector

import (
	"bytes"
)

func New(options ...Option) *SelectionVector {
	sv := &SelectionVector{
		Data: new(bytes.Buffer),
	}
	for _, option := range options {
		*sv = option(*sv)
	}
	return sv
}

type Option func(SelectionVector) SelectionVector

func WithSize(size int) Option {
	return func(sv SelectionVector) SelectionVector {
		sv.Data.Grow(size)
		return sv
	}
}
