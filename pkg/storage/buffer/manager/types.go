package manager

import (
	"sync"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/buffer/block/iface"
	"tae/pkg/storage/layout"
)

type EvictNode struct {
	Block iface.IBlockHandle
	Iter  types.IDX_T
}

type IEvictHolder interface {
	sync.Locker
	Enqueue(node *EvictNode)
	// Count() types.IDX_T
	Dequeue() *EvictNode
}

type SimpleEvictHolder struct {
	sync.Mutex
	Queue chan *EvictNode
}

const (
	SIMPLE_EVICT_HOLDER_CAPACITY = 100000
)

type BufferManager struct {
	sync.Mutex
	Blocks      map[layout.BlockId]iface.IBlockHandle // Manager is not responsible to Close handle
	TransientID layout.BlockId
	Pool        buf.IMemoryPool
	EvictHolder IEvictHolder
	// TempPath string
}
