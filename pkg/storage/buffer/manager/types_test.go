package manager

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"tae/pkg/storage/layout"
	"testing"
)

func TestManager(t *testing.T) {
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
	assert.True(t, ok)
	assert.True(t, mgr_h1.IsClosed())
	mgr_h2, ok := mgr.(*BufferManager).Blocks[blk_2]
	assert.True(t, ok)
	assert.False(t, mgr_h2.IsClosed())

	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 3)
	mgr.UnregisterBlock(blk_0, true)
	assert.Equal(t, len(mgr.(*BufferManager).Blocks), 2)
}
