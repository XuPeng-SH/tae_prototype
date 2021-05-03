package block

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
	"testing"
)

func TestBlock(t *testing.T) {
	blk_id := types.IDX_T(999)
	blk := NewBlockBuffer(blk_id)
	assert.Equal(t, blk.Capacity(), int64(layout.BLOCK_ALLOC_SIZE))
	assert.Equal(t, blk_id, blk.GetID())
	assert.Equal(t, buf.BLOCK_BUFFER, blk.GetType())
}
