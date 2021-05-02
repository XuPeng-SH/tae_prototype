package coldata

import (
	"tae/pkg/common/types"
	seg "tae/pkg/storage/table/segment"
)

type IColumnData interface {
}

type ColumnData struct {
	ColumnType types.LogicType
	ColumnIdx  types.IDX_T
	RowCount   types.IDX_T
	DataTree   seg.ISegmentTree
	UpdateTree seg.ISegmentTree
}
