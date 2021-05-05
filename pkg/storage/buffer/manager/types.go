package manager

import (
	"sync"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
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
}

type BufferManager struct {
	sync.Mutex
	Blocks      map[layout.BlockId]iface.IBlockHandle // Manager is not responsible to Close handle
	TransientID layout.BlockId
	Pool        buf.IMemoryPool
	// EvictHandle IEvictHandle
	// TempPath string
}
