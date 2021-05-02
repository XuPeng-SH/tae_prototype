package segment

import (
	"sync"
	"tae/pkg/common/types"
	// "tae/pkg/common/types/constants"
)

type ISegment interface {
	String() string
	ToString(verbose bool) string
	GetNext() ISegment
	Append(next ISegment)
	GetStartRow() types.IDX_T
	GetEndRow() types.IDX_T
	GetRowCount() types.IDX_T
	Capacity() types.IDX_T
}

const (
// SEGMENT_MAX_ROWS types.IDX_T = constants.STANDARD_VECTOR_SIZE * MOSEL_VECTOR_COUNT
)

type Segment struct {
	StartRow types.IDX_T
	RowCount types.IDX_T
	Next     ISegment
}

type ISegmentTree interface {
	// All interfaces are not thread-safe. Should call RLock or Lock manually
	String() string
	ToString(depth types.IDX_T) string
	GetRoot() ISegment
	GetTail() ISegment
	Depth() types.IDX_T
	WhichSeg(row types.IDX_T) ISegment
	WhichSegIdx(row types.IDX_T) types.IDX_T
	Append(new_seg ISegment)
	// ReferenceOther(other ISegmentTree)
}

type SegmentTree struct {
	sync.RWMutex
	Segments []ISegment
}
