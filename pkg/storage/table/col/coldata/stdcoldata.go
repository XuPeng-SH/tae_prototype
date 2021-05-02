package coldata

import (
	"tae/pkg/common/types"
)

type StdColumnData struct {
	IColumnData
	Validity IColumnData
}

func NewStdColumnData(col_type types.LogicType, col_idx types.IDX_T) IColumnData {
	data := &StdColumnData{
		IColumnData: NewColumnData(col_type, col_idx),
		Validity:    NewValidtyColumnData(col_idx),
	}
	return data
}
