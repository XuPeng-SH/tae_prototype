package validity_mask

import (
	"bytes"
	"tae/pkg/common/types"
)

func WhichEntry(idx types.IDX_T) types.IDX_T {
	return (idx + ((types.IDX_T)(BITS_PER_ENTRY) - 1)) / (types.IDX_T)(BITS_PER_ENTRY)
}

type Option func(ValidityMask) ValidityMask

func New(count int, options ...Option) *ValidityMask {
	vm := &ValidityMask{
		Data: new(bytes.Buffer),
	}
	vm.Data.Grow(count)
	for _, option := range options {
		*vm = option(*vm)
	}
	return vm
}

// func WithSize(size int) Option {
// 	return func(sv SelectionVector) SelectionVector {
// 		sv.Data.Grow(size)
// 		return sv
// 	}
// }
