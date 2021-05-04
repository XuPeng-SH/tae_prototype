package manager

import (
	"sync"
	"tae/pkg/common/types"
	"tae/pkg/storage/buffer/block/iface"
	"tae/pkg/storage/layout"
)

type EvictNode struct {
	Block    iface.IBlockHandle
	Sequence types.IDX_T
}

type IEvictHandle interface {
	Enqueue(node *EvictNode)
	Count() types.IDX_T
	// MakeSpace(gcc)
}

type BufferManager struct {
	sync.Mutex
	UsageSize   types.IDX_T
	Capacity    types.IDX_T
	Blocks      map[layout.BlockId]iface.IBlockHandle // Manager is not responsible to Close handle
	TransientID layout.BlockId
	// EvictHandle IEvictHandle
	// TempPath string
}
