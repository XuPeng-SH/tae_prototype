package block

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
	"testing"
)

func TestBlock(t *testing.T) {
	pool := buf.NewSimpleMemoryPool(layout.BLOCK_ALLOC_SIZE * 2)
	// node := pool.Get()
	blk_id := layout.NewBlockId(0, 0)
	node := pool.Get(layout.BLOCK_ALLOC_SIZE, false)
	blk := NewBlockBuffer(*blk_id, node)
	assert.Equal(t, blk.Capacity(), int64(layout.BLOCK_ALLOC_SIZE))
	assert.Equal(t, *blk_id, blk.GetID())
	assert.Equal(t, buf.BLOCK_BUFFER, blk.GetType())
	assert.Equal(t, types.IDX_T(blk.Capacity()), pool.GetUsageSize())

	blk0_1_id := layout.NewBlockId(0, 1)
	node2 := pool.Get(layout.BLOCK_ALLOC_SIZE, false)
	blk0_1 := NewBlockBuffer(*blk0_1_id, node2)
	assert.Equal(t, blk0_1.Capacity(), int64(layout.BLOCK_ALLOC_SIZE))
	assert.Equal(t, types.IDX_T(blk.Capacity()+blk0_1.Capacity()), pool.GetUsageSize())

	blk.Close()
	assert.Equal(t, blk.Capacity(), int64(0))
	assert.Equal(t, types.IDX_T(blk0_1.Capacity()), pool.GetUsageSize())

	blk0_1.Close()
	assert.Equal(t, blk0_1.Capacity(), int64(0))
	assert.Equal(t, types.IDX_0, pool.GetUsageSize())
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
