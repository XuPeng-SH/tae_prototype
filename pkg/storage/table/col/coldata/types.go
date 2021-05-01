package coldata

import (
	"tae/pkg/common/types"
)

type ColumnData struct {
	ColumnType types.LogicType
	ColumnIdx  types.SMIDX_T
	RowCount   types.IDX_T
}
