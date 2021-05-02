package coldata

import (
	"tae/pkg/common/types"
)

type IColumnData interface {
}

type ColumnData struct {
	ColumnType types.LogicType
	ColumnIdx  types.IDX_T
	RowCount   types.IDX_T
}
