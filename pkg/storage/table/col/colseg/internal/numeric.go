package internal

import "tae/pkg/common/types"

type NumericSegment struct {
	SimpleSegmentImpl
}

func NewNumericSegment(t types.PhysicalType, start_row types.IDX_T) *NumericSegment {
	seg := &NumericSegment{
		SimpleSegmentImpl: SimpleSegmentImpl{
			Type:     t,
			StartRow: start_row,
		},
	}
	return seg
}
