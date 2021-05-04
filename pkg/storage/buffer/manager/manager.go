package manager

import (
	"tae/pkg/common/types"
	blk "tae/pkg/storage/buffer/block"
	"tae/pkg/storage/layout"
)

var (
	_ IBufferManager = (*BufferManager)(nil)
)

func NewBufferManager(capacity types.IDX_T) IBufferManager {
	mgr := &BufferManager{
		Capacity:    capacity,
		TransientID: layout.MIN_TRANSIENT_BLOCK_ID,
		Blocks:      make(map[layout.BlockId]blk.IBlockHandle),
	}

	return mgr
}

func (mgr *BufferManager) RegisterBlock(blk_id layout.BlockId) blk.IBlockHandle {
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

func (mgr *BufferManager) UnregisterBlock(blk_id layout.BlockId, can_destroy bool) {
	if blk_id.IsTransientBlock() {
		// PXU TODO
		return
	}
	mgr.Lock()
	defer mgr.Unlock()
	delete(mgr.Blocks, blk_id)
}

func (mgr *BufferManager) Pin(handle blk.IBlockHandle) blk.IBufferHandle {
	handle.Lock()
	defer handle.Unlock()
	if handle.GetState() == blk.BLOCK_LOADED {
		// PXU TODO
		return nil
	}
	return nil
}
