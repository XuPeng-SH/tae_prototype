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
	}
	return handle
}

func (h *BlockHandle) Unload() {
	if h.State == BLOCK_UNLOAD {
		return
	}
	h.State = BLOCK_UNLOAD
}
