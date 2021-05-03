package buffer

import (
	"tae/pkg/common/types"
	"tae/pkg/storage/layout"
)

type IBlockBuffer interface {
	IBuffer
	GetID() types.IDX_T
}

type BlockBuffer struct {
	IBuffer
	ID types.IDX_T
}

var (
	_ IBuffer      = (*BlockBuffer)(nil)
	_ IBlockBuffer = (*BlockBuffer)(nil)
)

func NewBlockBuffer(id types.IDX_T) IBlockBuffer {
	bb := &BlockBuffer{
		IBuffer: NewBuffer(layout.BLOCK_ALLOC_SIZE),
		ID:      id,
	}
	bb.IBuffer.(*Buffer).Type = BLOCK_BUFFER
	return bb
}

func (bb *BlockBuffer) GetID() types.IDX_T {
	return bb.ID
}
