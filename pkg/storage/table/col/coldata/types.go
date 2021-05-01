package coldata

import (
	"tae/pkg/common/types"
)

type IColumnData interface {
}

type ColumnData struct {
	IColumnData
	ColumnType types.LogicType
	ColumnIdx  types.SMIDX_T
	RowCount   types.IDX_T
}
