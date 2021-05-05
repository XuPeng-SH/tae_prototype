package block

import (
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	blkif "tae/pkg/storage/buffer/block/iface"
	mgrif "tae/pkg/storage/buffer/manager/iface"
	"tae/pkg/storage/layout"
)

func NewBlockHandle(ctx *BlockHandleCtx) blkif.IBlockHandle {
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
	if blkif.AtomicCASState(&(h.State), blkif.BLOCK_LOADED, blkif.BLOCK_UNLOADING) {
		h.Buff.Close()
		h.Buff = nil
		blkif.AtomicStoreState(&(h.State), blkif.BLOCK_UNLOAD)
	}
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
	if !blkif.AtomicCASRTState(&(h.RTState), blkif.BLOCK_RT_RUNNING, blkif.BLOCK_RT_CLOSED) {
		// Cocurrent senario that other client already call Close before
		return nil
	}
	if h.Buff != nil {
		h.Buff.Close()
	}
	h.Manager.UnregisterBlock(h.ID, h.Destroyable)
	return nil
}

func (h *BlockHandle) IsClosed() bool {
	state := blkif.AtomicLoadRTState(&(h.RTState))
	return state == blkif.BLOCK_RT_CLOSED
}

func (h *BlockHandle) RollbackLoad() {
	if !blkif.AtomicCASState(&(h.State), blkif.BLOCK_LOADING, blkif.BLOCK_ROOLBACK) {
		return
	}
	h.UnRef()
	if h.Buff != nil {
		h.Buff.Close()
	}
	h.Buff = nil
	blkif.AtomicStoreState(&(h.State), blkif.BLOCK_UNLOAD)
}

func (h *BlockHandle) PrepareLoad() bool {
	return blkif.AtomicCASState(&(h.State), blkif.BLOCK_UNLOAD, blkif.BLOCK_LOADING)
}

func (h *BlockHandle) CommitLoad() error {
	if !blkif.AtomicCASState(&(h.State), blkif.BLOCK_LOADING, blkif.BLOCK_COMMIT) {
		return types.ErrLogicError
	}

	// TODO: Load content from io here

	if !blkif.AtomicCASState(&(h.State), blkif.BLOCK_COMMIT, blkif.BLOCK_LOADED) {
		return types.ErrLogicError
	}
	return nil
}

func (h *BlockHandle) MakeHandle() blkif.IBufferHandle {
	if blkif.AtomicLoadState(&(h.State)) != blkif.BLOCK_LOADED {
		panic("Should not call MakeHandle not BLOCK_LOADED")
	}
	return NewBufferHandle(h, h.Manager)
}

func (h *BlockHandle) SetBuffer(buf buf.IBuffer) error {
	if h.Buff != nil || h.Capacity != types.IDX_T(buf.Capacity()) {
		return types.ErrLogicError
	}
	h.Buff = buf
	return nil
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
