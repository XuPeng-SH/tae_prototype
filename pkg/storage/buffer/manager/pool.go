package manager

import (
	"tae/pkg/common/types"
)

func NewSimpleMemoryPool(capacity types.IDX_T) IMemoryPool {
	pool := &SimpleMemoryPool{
		Capacity: capacity,
	}
	return pool
}

func (pool *SimpleMemoryPool) GetCapacity() types.IDX_T {
	return types.AtomicLoad(&(pool.Capacity))
}

func (pool *SimpleMemoryPool) SetCapacity(capacity types.IDX_T) {
	types.AtomicStore(&(pool.Capacity), capacity)
}

func (pool *SimpleMemoryPool) GetUsageSize() types.IDX_T {
	return types.AtomicLoad(&(pool.UsageSize))
}

// Only for temp test
func (pool *SimpleMemoryPool) Get(size types.IDX_T) (node *PoolNode) {
	capacity := types.AtomicLoad(&(pool.Capacity))
	currsize := types.AtomicLoad(&(pool.UsageSize))
	postsize := size + currsize
	if postsize > capacity {
		return &PoolNode{Buff: []byte{}, Pool: pool}
	}
	for !types.AtomicCAS(&(pool.UsageSize), currsize, postsize) {
		currsize = types.AtomicLoad(&(pool.UsageSize))
		postsize += currsize + size
		if postsize > capacity {
			return &PoolNode{Buff: []byte{}, Pool: pool}
		}
	}
	buf := make([]byte, size)
	return &PoolNode{Buff: buf, Pool: pool}
}

// Only for temp test
func (pool *SimpleMemoryPool) Put(node *PoolNode) {
	size := int64(len(node.Buff))
	if size == 0 {
		return
	}
	usagesize := types.AtomicAdd(&(pool.UsageSize), int64(size))
	if usagesize > pool.Capacity {
		panic("")
	}
}
