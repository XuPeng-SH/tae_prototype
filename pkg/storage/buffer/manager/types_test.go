package manager

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	blkif "tae/pkg/storage/buffer/block/iface"
	"tae/pkg/storage/layout"
	"testing"
)

func TestManagerBasic(t *testing.T) {
	mgr := NewBufferManager(types.IDX_T(1))
	blk_0 := layout.BlockId{Part: uint32(0), Offset: uint32(0)}
	blk_1 := layout.BlockId{Part: uint32(0), Offset: uint32(1)}
	blk_01 := layout.BlockId{Part: uint32(0), Offset: uint32(0)}
	blk_2 := layout.BlockId{Part: uint32(0), Offset: uint32(2)}
	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 0)
	h0 := mgr.RegisterBlock(blk_0)
	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 1)
	assert.Equal(t, blk_0, h0.GetID())
	h1 := mgr.RegisterBlock(blk_1)
	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 2)
	assert.Equal(t, blk_1, h1.GetID())

	h01 := mgr.RegisterBlock(blk_01)
	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 2)
	assert.Equal(t, blk_01, h01.GetID())
	h2 := mgr.RegisterBlock(blk_2)
	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 3)
	assert.Equal(t, blk_2, h2.GetID())

	h1.Close()
	assert.True(t, h1.IsClosed())
	mgr_h1, ok := mgr.(*BufferManager).Blocks[blk_1]
	assert.False(t, ok)
	assert.Equal(t, mgr_h1, nil)
	mgr_h2, ok := mgr.(*BufferManager).Blocks[blk_2]
	assert.True(t, ok)
	assert.False(t, mgr_h2.IsClosed())

	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 2)
	mgr.UnregisterBlock(blk_0, true)
	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 1)
}

func TestManager2(t *testing.T) {
	mgr := NewBufferManager(types.IDX_T(1024))

	blk_0_0 := layout.NewBlockId(0, 0)
	h_0_0 := mgr.RegisterBlock(*blk_0_0)
	assert.Equal(t, h_0_0.GetID(), *blk_0_0)
	assert.False(t, h_0_0.HasRef())
	b := mgr.Pin(h_0_0)
	assert.Equal(t, b, nil)
	new_cap := h_0_0.GetCapacity() * 2
	mgr.SetCapacity(new_cap)
	assert.Equal(t, mgr.GetCapacity(), new_cap)
	t.Log(new_cap)

	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 1)
	assert.False(t, h_0_0.HasRef())

	b_0_0 := mgr.Pin(h_0_0)
	assert.Equal(t, b_0_0.GetID(), *blk_0_0)
	assert.True(t, h_0_0.HasRef())
	b_0_0.Close()
	assert.False(t, h_0_0.HasRef())
	b_0_0_0 := mgr.Pin(h_0_0)
	assert.True(t, h_0_0.HasRef())
	b_0_0_1 := mgr.Pin(h_0_0)
	assert.True(t, h_0_0.HasRef())
	b_0_0_2 := mgr.Pin(h_0_0)
	assert.True(t, h_0_0.HasRef())
	b_0_0_0.Close()
	assert.True(t, h_0_0.HasRef())
	b_0_0_2.Close()
	assert.True(t, h_0_0.HasRef())
	b_0_0_1.Close()
	assert.False(t, h_0_0.HasRef())
}

func TestManager3(t *testing.T) {
	capacity := layout.BLOCK_ALLOC_SIZE * 100
	mgr := NewBufferManager(capacity)
	assert.Equal(t, mgr.GetCapacity(), capacity)

	blk_id_0_0 := *layout.NewBlockId(0, 0)
	h_0_0 := mgr.RegisterBlock(blk_id_0_0)
	assert.True(t, h_0_0 != nil)
	assert.Equal(t, h_0_0.GetID(), blk_id_0_0)
	assert.Equal(t, h_0_0.GetState(), blkif.BLOCK_UNLOAD)
	assert.Equal(t, mgr.GetCapacity(), capacity)

	// bh_0_0 := mgr.Pin(h_0_0)

}
