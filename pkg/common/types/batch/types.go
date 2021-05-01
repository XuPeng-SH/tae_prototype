package batch

import (
	"tae/pkg/common/types"
	"tae/pkg/common/types/value"
	"tae/pkg/common/types/vector"
)

type IBatch interface {
	Cols() types.IDX_T
	Rows() types.IDX_T
	GetCell(row, col types.IDX_T) value.Value
	SetCell(row, col types.IDX_T, val value.Value)
	Append(other IBatch)
	// Reset()
	// Drop()
	// ReferenceOther(other IBatch)
	GetTypes() []types.LogicType
	// String() string
}

type Batch struct {
	Data     []*vector.Vector
	RowCount types.IDX_T
}
