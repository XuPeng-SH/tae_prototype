package table

import (
	"fmt"
	"sync"
	"tae/pkg/common/types"
	"tae/pkg/storage/table/col/coldata"
	seg "tae/pkg/storage/table/segment"
)

type PhysicalTableState uint8

const (
	ACTIVE PhysicalTableState = iota
	STALE
)

func (state PhysicalTableState) String() string {
	switch state {
	case ACTIVE:
		return "ACTIVE"
	case STALE:
		return "STALE"
	}
	panic(fmt.Sprintf("UNKNOWN state: %d", state))
}

type IPhysicalTable interface {
	GetRowCount() types.IDX_T
}

type PhysicalTable struct {
	IPhysicalTable
	Mu          sync.Mutex
	RowCount    types.IDX_T
	Versions    seg.ISegmentTree
	Columns     []coldata.IColumnData
	Fingerprint types.IDX_T
	State       PhysicalTableState
}
