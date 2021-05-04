package manager

import (
	"fmt"
	"tae/pkg/common/types"
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
		Capacity:    capacity,
		TransientID: layout.MIN_TRANSIENT_BLOCK_ID,
		Blocks:      make(map[layout.BlockId]blkif.IBlockHandle),
	}

	return mgr
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
		ID: blk_id,
	}
	handle = blk.NewBlockHandle(&ctx)
	mgr.Blocks[blk_id] = handle
	return handle
}

func (mgr *BufferManager) GetUsageSize() types.IDX_T {
	return types.AtomicLoad(&(mgr.UsageSize))
}

func (mgr *BufferManager) GetCapacity() types.IDX_T {
	return types.AtomicLoad(&(mgr.Capacity))
}

func (mgr *BufferManager) SetCapacity(capacity types.IDX_T) {
	mgr.Lock()
	defer mgr.Unlock()
	if !mgr.makeSpace(0, capacity) {
		panic(fmt.Sprintf("Cannot makeSpace(%d,%d)", 0, capacity))
	}
	types.AtomicStore(&(mgr.Capacity), capacity)
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
	// TODO
	return true
}

func (mgr *BufferManager) Pin(handle blkif.IBlockHandle) blkif.IBufferHandle {
	handle.Lock()
	defer handle.Unlock()
	if handle.GetState() == blkif.BLOCK_LOADED {
		// PXU TODO
		return nil
	}
	if mgr.makeSpace(handle.GetCapacity(), mgr.GetCapacity()) {
		panic(fmt.Sprintf("Cannot makeSpace(%d,%d)", handle.GetCapacity(), mgr.GetCapacity()))
	}
	handle.Ref()
	return nil
}
