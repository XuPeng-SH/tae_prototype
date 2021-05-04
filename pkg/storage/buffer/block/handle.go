package block

import (
	"tae/pkg/common/types"
	"tae/pkg/storage/buffer/block/iface"
	"tae/pkg/storage/layout"
)

func NewBlockHandle(ctx *BlockHandleCtx) iface.IBlockHandle {
	size := layout.BLOCK_ALLOC_SIZE
	state := iface.BLOCK_UNLOAD
	if ctx.Buff != nil {
		size = types.IDX_T(ctx.Buff.Capacity())
		state = iface.BLOCK_UNLOAD
	}
	handle := &BlockHandle{
		ID:       ctx.ID,
		Buff:     ctx.Buff,
		Capacity: size,
		State:    state,
		RTState:  iface.BLOCK_RT_RUNNING,
	}
	return handle
}

func (h *BlockHandle) Unload() {
	if h.State == iface.BLOCK_UNLOAD {
		return
	}
	h.State = iface.BLOCK_UNLOAD
}

func (h *BlockHandle) GetCapacity() types.IDX_T {
	return h.Capacity
}

func (h *BlockHandle) Ref() {
	types.AtomicAdd(&(h.Refs), 1)
}

func (h *BlockHandle) UnRef() {
	old := types.AtomicLoad(&(h.Refs))
	if old == types.IDX_0 {
		return
	}
	types.AtomicCAS(&(h.Refs), old, old-1)
}

func (h *BlockHandle) HasRef() bool {
	v := types.AtomicLoad(&(h.Refs))
	return v > types.IDX_0
}

func (h *BlockHandle) GetID() layout.BlockId {
	return h.ID
}

func (h *BlockHandle) GetState() iface.BlockState {
	return h.State
}

func (h *BlockHandle) Close() error {
	h.RTState = iface.BLOCK_RT_CLOSED
	return nil
}

// PXU TODO
func (h *BlockHandle) IsClosed() bool {
	return h.RTState == iface.BLOCK_RT_CLOSED
}

func (h *BlockHandle) Load() iface.IBufferHandle {
	if h.State == iface.BLOCK_LOADED {
		return nil
	}
	// TODO
	return nil
}
