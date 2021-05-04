package layout

import (
	// "fmt"
	"tae/pkg/common/types"
)

var (
	BLOCK_SECTOR_SIZE   types.IDX_T = 4 * K
	BLOCK_HEAD_SIZE     types.IDX_T = 32
	BLOCK_ALLOC_SIZE    types.IDX_T = 256 * K
	BLOCK_DATA_SIZE                 = BLOCK_ALLOC_SIZE - BLOCK_HEAD_SIZE
	BLOCK_CHECKSUM_SIZE             = BLOCK_HEAD_SIZE
)

func NewBlockId(part, offset uint32) *BlockId {
	blk_id := &BlockId{
		Part:   part,
		Offset: offset,
	}
	return blk_id
}

func (bid BlockId) IsTransientBlock() bool {
	return bid.Part >= MIN_TRANSIENT_BLOCK_ID.Part
}
