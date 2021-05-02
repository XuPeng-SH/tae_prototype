package coldata

import (
	"fmt"
)

func (cdata *ColumnData) String() string {
	return fmt.Sprintf("CData(%s,%d,%d)", cdata.ColumnType.String(), cdata.ColumnIdx, cdata.RowCount)
}
