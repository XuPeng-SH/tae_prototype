package buffer

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

func (pool *SimpleMemoryPool) SetCapacity(capacity types.IDX_T) error {
	if capacity < types.AtomicLoad(&(pool.Capacity)) {
		return types.ErrLogicError
	}
	types.AtomicStore(&(pool.Capacity), capacity)
	return nil
}

func (pool *SimpleMemoryPool) GetUsageSize() types.IDX_T {
	return types.AtomicLoad(&(pool.UsageSize))
}

// Only for temp test
func (pool *SimpleMemoryPool) Get(size types.IDX_T, lazy bool) (node *PoolNode) {
	capacity := types.AtomicLoad(&(pool.Capacity))
	currsize := types.AtomicLoad(&(pool.UsageSize))
	postsize := size + currsize
	if postsize > capacity {
		return nil
		// return &PoolNode{Data: []byte{}, Pool: pool}
	}
	for !types.AtomicCAS(&(pool.UsageSize), currsize, postsize) {
		currsize = types.AtomicLoad(&(pool.UsageSize))
		postsize += currsize + size
		if postsize > capacity {
			return nil
			// return &PoolNode{Data: []byte{}, Pool: pool}
		}
	}
	buf := []byte{}
	if !lazy {
		buf = make([]byte, size)
	}
	return &PoolNode{Data: buf, Pool: pool, Size: size}
}

// Only for temp test
func (pool *SimpleMemoryPool) DoAlloc(node *PoolNode) {
	if len(node.Data) == int(node.Size) {
		return
	}
	node.Data = make([]byte, node.Size)
}

// Only for temp test
func (pool *SimpleMemoryPool) Put(node *PoolNode) {
	size := int64(len(node.Data))
	if size == 0 {
		return
	}
	usagesize := types.AtomicAdd(&(pool.UsageSize), -1*int64(size))
	if usagesize > pool.Capacity {
		panic("")
	}
	node.Data = nil
}
