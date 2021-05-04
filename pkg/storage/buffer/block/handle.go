package block

import (
	"tae/pkg/common/types"
	"tae/pkg/storage/layout"
)

func NewBlockHandle(ctx *BlockHandleCtx) IBlockHandle {
	size := layout.BLOCK_ALLOC_SIZE
	state := BLOCK_UNLOAD
	if ctx.Buff != nil {
		size = types.IDX_T(ctx.Buff.Capacity())
		state = BLOCK_UNLOAD
	}
	handle := &BlockHandle{
		ID:       ctx.ID,
		Buff:     ctx.Buff,
		Capacity: size,
		State:    state,
		RTState:  BLOCK_RT_RUNNING,
	}
	return handle
}

func (h *BlockHandle) Unload() {
	if h.State == BLOCK_UNLOAD {
		return
	}
	h.State = BLOCK_UNLOAD
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

func (h *BlockHandle) GetState() BlockState {
	return h.State
}

func (h *BlockHandle) Close() error {
	h.RTState = BLOCK_RT_CLOSED
	return nil
}

// PXU TODO
func (h *BlockHandle) IsClosed() bool {
	return h.RTState == BLOCK_RT_CLOSED
}
