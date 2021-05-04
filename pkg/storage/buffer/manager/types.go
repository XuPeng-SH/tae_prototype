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
}

type PoolNode struct {
	Buff []byte
	Pool IMemoryPool
}

type IMemoryPool interface {
	Get(size types.IDX_T) (node *PoolNode)
	Put(node *PoolNode)
	GetCapacity() types.IDX_T
	SetCapacity(capacity types.IDX_T)
	GetUsageSize() types.IDX_T
}

type SimpleMemoryPool struct {
	Capacity  types.IDX_T
	UsageSize types.IDX_T
}

type BufferManager struct {
	sync.Mutex
	// UsageSize   types.IDX_T
	// Capacity    types.IDX_T
	Blocks      map[layout.BlockId]iface.IBlockHandle // Manager is not responsible to Close handle
	TransientID layout.BlockId
	Pool        IMemoryPool
	// EvictHandle IEvictHandle
	// TempPath string
}
