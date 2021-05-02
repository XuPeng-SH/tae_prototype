package segment

import (
	"fmt"
	"tae/pkg/common/types"
)

var (
	_ ISegment = (*Segment)(nil)
)

func NewSegment(start, count types.IDX_T) ISegment {
	seg := &Segment{
		StartRow: start,
		RowCount: count,
	}
	return seg
}

func (seg *Segment) GetNext() ISegment {
	return seg.Next
}

func (seg *Segment) GetStartRow() types.IDX_T {
	return seg.StartRow
}

func (seg *Segment) GetEndRow() types.IDX_T {
	return seg.StartRow + seg.RowCount
}

func (seg *Segment) GetRowCount() types.IDX_T {
	return seg.RowCount
}

func (seg *Segment) Capacity() types.IDX_T {
	// PXU TODO
	return seg.RowCount
}

func (seg *Segment) Append(next ISegment) {
	seg.Next = next
}

func (seg *Segment) String() string {
	return seg.ToString(true)
}

func (seg *Segment) ToString(verbose bool) string {
	if verbose {
		return fmt.Sprintf("Seg(%v, %v)[HasNext:%v]", seg.StartRow, seg.RowCount, seg.Next != nil)
	}
	return fmt.Sprintf("(%v, %v)", seg.StartRow, seg.RowCount)
}
