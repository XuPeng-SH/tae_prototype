package coldata

import (
	"fmt"
	"tae/pkg/common/types"
)

func NewColumnData(col_type types.LogicType, col_idx types.IDX_T) IColumnData {
	data := &ColumnData{
		ColumnType: col_type,
		ColumnIdx:  col_idx,
	}
	return data
}

func (cdata *ColumnData) String() string {
	return fmt.Sprintf("CData(%s,%d,%d)", cdata.ColumnType.String(), cdata.ColumnIdx, cdata.RowCount)
}
