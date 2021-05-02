package batch

import (
	"tae/pkg/common/types"
	"tae/pkg/common/types/value"
	// "tae/pkg/common/types/vector"
	"fmt"
)

var (
	_ IBatch = (*Batch)(nil)
)

func NewBatch() *Batch {
	bat := &Batch{}
	return bat
}

func (bat *Batch) Cols() types.IDX_T {
	return (types.IDX_T)(len(bat.Data))
}

func (bat *Batch) Rows() types.IDX_T {
	return bat.RowCount
}

func (bat *Batch) GetCell(row, col types.IDX_T) *value.Value {
	if row >= bat.RowCount || col >= (types.IDX_T)(len(bat.Data)) {
		panic(fmt.Sprintf("(row,col)=(%v,%v) is out of range", row, col))
	}
	return bat.Data[col].GetValue(row).(*value.Value)
}

func (bat *Batch) Verify() {
}

func (bat *Batch) Flatten() {
	for _, vec := range bat.Data {
		vec.Flatten(bat.RowCount)
	}
}

func (bat *Batch) SetCell(row, col types.IDX_T, val *value.Value) {
	if row >= bat.RowCount || col >= (types.IDX_T)(len(bat.Data)) {
		panic(fmt.Sprintf("(row,col)=(%v,%v) is out of range", row, col))
	}
	bat.Data[col].SetValue(row, val)
}

func (bat *Batch) Append(other IBatch) {
	if other.Rows() == 0 {
		return
	}
	if other.Cols() != (types.IDX_T)(len(bat.Data)) {
		panic(fmt.Sprintf("Column count mismatch"))
	}

}

func (bat *Batch) GetTypes() []types.LogicType {
	ts := make([]types.LogicType, 0, len(bat.Data))
	for _, col := range bat.Data {
		ts = append(ts, col.GetLogicType())
	}
	return ts
}
