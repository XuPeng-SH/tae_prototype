package updateseg

import (
	"sync"
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	// "tae/pkg/common/types/selvec"
	"tae/pkg/storage/table/col/coldata"
	seg "tae/pkg/storage/table/segment"
)

const (
	MOSEL_VECTOR_COUNT      types.IDX_T = 100
	MOSEL_SIZE                          = constants.STANDARD_VECTOR_SIZE * MOSEL_VECTOR_COUNT
	UPDATE_SEGMENT_MAX_ROWS             = constants.STANDARD_VECTOR_SIZE * MOSEL_VECTOR_COUNT
)

type IUpdateSegment interface {
	HasUpdate() bool
	HasVectorUpdate(vec_idx types.IDX_T) bool
	HasVectorRangeUpdate(vstart_idx, vend_idx types.IDX_T) bool
	FindSegByVecIdx(vec_idx types.IDX_T) seg.ISegment
	FindSegByRowIdx(row_idx types.IDX_T) seg.ISegment
}

type UpdateNode struct {
	// SelData selvec.SelectionData
}

type UpdateTree struct {
	Nodes [MOSEL_VECTOR_COUNT]UpdateNode
}

type UpdateSegment struct {
	seg.ISegment
	Data *coldata.ColumnData
	Mu   sync.RWMutex
	Tree *UpdateTree
}
