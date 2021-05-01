package segment

import (
	"sync"
	"tae/pkg/common/types"
)

type ISegment interface {
	String() string
	ToString(verbose bool) string
	GetNext() ISegment
	Append(ISegment)
	GetStartRow() types.IDX_T
	GetEndRow() types.IDX_T
	GetRowCount() types.IDX_T
}

type Segment struct {
	ISegment
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
