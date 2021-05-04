package block

import (
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/buffer/block/iface"
	"tae/pkg/storage/layout"
)

var (
	_ buf.IBuffer        = (*BlockBuffer)(nil)
	_ iface.IBlockBuffer = (*BlockBuffer)(nil)
)

func NewBlockBuffer(id layout.BlockId) iface.IBlockBuffer {
	bb := &BlockBuffer{
		IBuffer: buf.NewBuffer(layout.BLOCK_ALLOC_SIZE),
		ID:      id,
	}
	bb.IBuffer.(*buf.Buffer).Type = buf.BLOCK_BUFFER
	return bb
}

func (bb *BlockBuffer) GetID() layout.BlockId {
	return bb.ID
}
