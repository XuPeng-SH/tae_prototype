package internal

import (
	"tae/pkg/common/types"
)

type SimpleSegmentImpl struct {
	Type           types.PhysicalType
	VectorSize     types.IDX_T
	VectorCapacity types.IDX_T
	Count          types.IDX_T
	StartRow       types.IDX_T
}
