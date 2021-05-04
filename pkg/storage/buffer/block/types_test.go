package block

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
	"testing"
)

func TestBlock(t *testing.T) {
	blk_id := layout.BlockId{Part: uint32(0), Offset: uint32(0)}
	blk := NewBlockBuffer(blk_id)
	assert.Equal(t, blk.Capacity(), int64(layout.BLOCK_ALLOC_SIZE))
	assert.Equal(t, blk_id, blk.GetID())
	assert.Equal(t, buf.BLOCK_BUFFER, blk.GetType())
}

func TestHandle(t *testing.T) {
	blk_0 := layout.BlockId{Part: uint32(0), Offset: uint32(0)}
	ctx := BlockHandleCtx{
		ID: blk_0,
	}
	handle := NewBlockHandle(&ctx)
	assert.Equal(t, handle.(*BlockHandle).Refs, types.IDX_0)
	assert.False(t, handle.HasRef())
	handle.Ref()
	assert.Equal(t, handle.(*BlockHandle).Refs, types.IDX_1)
	assert.True(t, handle.HasRef())
	handle.Ref()
	assert.Equal(t, handle.(*BlockHandle).Refs, types.IDX_2)
	assert.True(t, handle.HasRef())
	handle.UnRef()
	assert.Equal(t, handle.(*BlockHandle).Refs, types.IDX_1)
	assert.True(t, handle.HasRef())
	handle.UnRef()
	assert.Equal(t, handle.(*BlockHandle).Refs, types.IDX_0)
	assert.False(t, handle.HasRef())
	handle.UnRef()
	assert.Equal(t, handle.(*BlockHandle).Refs, types.IDX_0)
	assert.False(t, handle.HasRef())
	t.Log(handle.(*BlockHandle).Refs)
}
