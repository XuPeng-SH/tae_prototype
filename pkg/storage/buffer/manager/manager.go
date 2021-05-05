package manager

import (
	log "github.com/sirupsen/logrus"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	blk "tae/pkg/storage/buffer/block"
	blkif "tae/pkg/storage/buffer/block/iface"
	mgrif "tae/pkg/storage/buffer/manager/iface"
	"tae/pkg/storage/layout"
)

var (
	_ mgrif.IBufferManager = (*BufferManager)(nil)
)

func NewBufferManager(capacity types.IDX_T) mgrif.IBufferManager {

	mgr := &BufferManager{
		Pool:        buf.NewSimpleMemoryPool(capacity),
		TransientID: layout.MIN_TRANSIENT_BLOCK_ID,
		Blocks:      make(map[layout.BlockId]blkif.IBlockHandle),
	}

	return mgr
}

func (mgr *BufferManager) GetPool() buf.IMemoryPool {
	return mgr.Pool
}

func (mgr *BufferManager) RegisterBlock(blk_id layout.BlockId) blkif.IBlockHandle {
	mgr.Lock()
	defer mgr.Unlock()

	handle, ok := mgr.Blocks[blk_id]
	if ok {
		if !handle.IsClosed() {
			return handle
		}
	}
	ctx := blk.BlockHandleCtx{
		ID:      blk_id,
		Manager: mgr,
	}
	handle = blk.NewBlockHandle(&ctx)
	mgr.Blocks[blk_id] = handle
	return handle
}

func (mgr *BufferManager) GetUsageSize() types.IDX_T {
	return mgr.Pool.GetUsageSize()
}

func (mgr *BufferManager) GetCapacity() types.IDX_T {
	return mgr.Pool.GetCapacity()
}

// Temp only can SetCapacity with larger size
func (mgr *BufferManager) SetCapacity(capacity types.IDX_T) error {
	mgr.Lock()
	defer mgr.Unlock()
	// if !mgr.makeSpace(0, capacity) {
	// 	panic(fmt.Sprintf("Cannot makeSpace(%d,%d)", 0, capacity))
	// }
	// types.AtomicStore(&(mgr.Capacity), capacity)
	return mgr.Pool.SetCapacity(capacity)
}

func (mgr *BufferManager) UnregisterBlock(blk_id layout.BlockId, can_destroy bool) {
	if blk_id.IsTransientBlock() {
		// PXU TODO
		return
	}
	mgr.Lock()
	defer mgr.Unlock()
	delete(mgr.Blocks, blk_id)
}

func (mgr *BufferManager) makeSpace(free_size, upper_limit types.IDX_T) bool {

	// for !types.AtomicCAS(&(mgr.UsageSize), currsize, postsize) {
	// 	currsize = types.AtomicLoad(&(pool.UsageSize))
	// 	postsize += currsize + size
	// 	if postsize > capacity {
	// 		return nil
	// 		// return &PoolNode{Data: []byte{}, Pool: pool}
	// 	}
	// }
	// for
	if free_size > upper_limit {
		return false
	}
	// TODO
	return true
}

func (mgr *BufferManager) Unpin(handle blkif.IBlockHandle) {
	mgr.Lock()
	defer mgr.Unlock()
	if !handle.UnRef() {
		panic("logic error")
	}
	if !handle.HasRef() {
		// Mark handle as stale
		// Temp to delete the handle from map
		// FIXME
		// delete(mgr.Blocks, handle.GetID())
	}
}

func (mgr *BufferManager) makePoolNode(capacity types.IDX_T) *buf.PoolNode {
	node := mgr.Pool.Get(capacity, false)
	for node == nil {
		// TODO
		return nil
	}
	return node
}

// TODO: Make Pin lock-free
func (mgr *BufferManager) Pin(handle blkif.IBlockHandle) blkif.IBufferHandle {
	handle.Lock()
	defer handle.Unlock()
	if handle.PrepareLoad() {
		node := mgr.makePoolNode(handle.GetCapacity())
		if node == nil {
			handle.RollbackLoad()
			log.Warnf("Cannot makeSpace(%d,%d)", handle.GetCapacity(), mgr.GetCapacity())
			return nil
		}
		buf := blk.NewBlockBuffer(handle.GetID(), node)
		handle.SetBuffer(buf)
		if err := handle.CommitLoad(); err != nil {
			handle.RollbackLoad()
			panic(err.Error())
		}
	}
	handle.Ref()
	return handle.MakeHandle()
}
