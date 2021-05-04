package block

import (
	"tae/pkg/common/types"
	blkif "tae/pkg/storage/buffer/block/iface"
	mgrif "tae/pkg/storage/buffer/manager/iface"
	"tae/pkg/storage/layout"
)

func NewBlockHandle(ctx *BlockHandleCtx) blkif.IBlockHandle {
	// if ctx.ID == nil {
	// 	panic(fmt.Sprintf("Block id should be specified"))
	// }
	size := layout.BLOCK_ALLOC_SIZE
	state := blkif.BLOCK_UNLOAD
	if ctx.Buff != nil {
		size = types.IDX_T(ctx.Buff.Capacity())
		state = blkif.BLOCK_UNLOAD
	}
	handle := &BlockHandle{
		ID:       ctx.ID,
		Buff:     ctx.Buff,
		Capacity: size,
		State:    state,
		RTState:  blkif.BLOCK_RT_RUNNING,
		Manager:  ctx.Manager,
	}
	return handle
}

func (h *BlockHandle) Unload() {
	if h.State == blkif.BLOCK_UNLOAD {
		return
	}
	h.State = blkif.BLOCK_UNLOAD
}

func (h *BlockHandle) GetCapacity() types.IDX_T {
	return h.Capacity
}

func (h *BlockHandle) Ref() {
	types.AtomicAdd(&(h.Refs), 1)
}

func (h *BlockHandle) UnRef() bool {
	old := types.AtomicLoad(&(h.Refs))
	if old == types.IDX_0 {
		return false
	}
	return types.AtomicCAS(&(h.Refs), old, old-1)
}

func (h *BlockHandle) HasRef() bool {
	v := types.AtomicLoad(&(h.Refs))
	return v > types.IDX_0
}

func (h *BlockHandle) GetID() layout.BlockId {
	return h.ID
}

func (h *BlockHandle) GetState() blkif.BlockState {
	return h.State
}

func (h *BlockHandle) Close() error {
	h.RTState = blkif.BLOCK_RT_CLOSED
	return nil
}

func (h *BlockHandle) IsClosed() bool {
	// TODO
	return h.RTState == blkif.BLOCK_RT_CLOSED
}

func (h *BlockHandle) Load() blkif.IBufferHandle {
	if h.State == blkif.BLOCK_LOADED {
		return NewBufferHandle(h, h.Manager)
	}

	// TODO
	h.State = blkif.BLOCK_LOADED
	return NewBufferHandle(h, h.Manager)
}

func NewBufferHandle(blk blkif.IBlockHandle, mgr mgrif.IBufferManager) blkif.IBufferHandle {
	h := &BufferHandle{
		Handle:  blk,
		Manager: mgr,
	}
	return h
}

func (h *BufferHandle) GetID() layout.BlockId {
	return h.Handle.GetID()
}

func (h *BufferHandle) Close() error {
	h.Manager.Unpin(h.Handle)
	return nil
}
