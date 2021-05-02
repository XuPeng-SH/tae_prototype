package batch

import (
	"tae/pkg/common/types"
)

func NewCollection() *Collection {
	return &Collection{}
}

func (coll *Collection) Cols() types.IDX_T {
	return types.IDX_T(len(coll.Types))
}

func (coll *Collection) GetCount() types.IDX_T {
	return coll.Count
}

func (coll *Collection) AppendBatch(bat IBatch) {
	if bat.Rows() == 0 {
		return
	}
	bat.Verify()
	coll.Count += bat.Rows()

	// offset := types.IDX_0
	// remaining := bat.Rows()
	if len(coll.Types) == 0 {
		coll.Types = bat.GetTypes()
	} else {
		if coll.Cols() != bat.Cols() {
			panic("mismatch types count")
		}
		target_col_types := bat.GetTypes()
		for col_idx, col_type := range coll.Types {
			if col_type != target_col_types[col_idx] {
				panic("mismatch type")
			}
		}
	}
}

func (coll *Collection) Append(other *Collection) {
	for _, bat := range other.Data {
		coll.AppendBatch(bat)
	}
}
