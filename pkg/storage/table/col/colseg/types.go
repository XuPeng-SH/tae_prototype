package colseg

import (
	"tae/pkg/common/types"
	base "tae/pkg/storage/table/segment"
)

type ColumnSegmentType uint8

const (
	TRANSIENT ColumnSegmentType = iota
	PERSISTENT
)

type IColumnSegment interface {
}

type ColumnSegment struct {
	IColumnSegment
	base.Segment
	ColumnType types.LogicType
	Type       ColumnSegmentType
}
