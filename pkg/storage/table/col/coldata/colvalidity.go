package coldata

import (
	"tae/pkg/common/types"
)

type ValidityColumnData struct {
	IColumnData
}

func NewValidtyColumnData(col_idx types.IDX_T) IColumnData {
	data := &ValidityColumnData{
		IColumnData: NewColumnData(types.LT_VALIDITY, col_idx),
	}
	return data
}
