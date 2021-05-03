package block

import (
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
)

var (
	_ buf.IBuffer  = (*BlockBuffer)(nil)
	_ IBlockBuffer = (*BlockBuffer)(nil)
)

func NewBlockBuffer(id types.IDX_T) IBlockBuffer {
	bb := &BlockBuffer{
		IBuffer: buf.NewBuffer(layout.BLOCK_ALLOC_SIZE),
		ID:      id,
	}
	bb.IBuffer.(*buf.Buffer).Type = buf.BLOCK_BUFFER
	return bb
}

func (bb *BlockBuffer) GetID() types.IDX_T {
	return bb.ID
}
