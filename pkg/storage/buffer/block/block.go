package block

import (
	log "github.com/sirupsen/logrus"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/buffer/block/iface"
	"tae/pkg/storage/layout"
)

var (
	_ buf.IBuffer        = (*BlockBuffer)(nil)
	_ iface.IBlockBuffer = (*BlockBuffer)(nil)
)

func NewBlockBuffer(id layout.BlockId, node *buf.PoolNode) iface.IBlockBuffer {
	if node == nil || node.Size != layout.BLOCK_ALLOC_SIZE {
		log.Warnf("NewBlockBuffer should accept node of size: %d", layout.BLOCK_ALLOC_SIZE)
		return nil
	}
	ibuf := buf.NewBuffer(node)
	bb := &BlockBuffer{
		IBuffer: ibuf,
		ID:      id,
	}
	bb.IBuffer.(*buf.Buffer).Type = buf.BLOCK_BUFFER
	return bb
}

func (bb *BlockBuffer) GetID() layout.BlockId {
	return bb.ID
}
