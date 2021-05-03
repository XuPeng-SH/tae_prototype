package layout

import (
	// "fmt"
	"tae/pkg/common/types"
)

const (
	BLOCK_INVALID                   = types.IDX_MAX
	BLOCK_MAX                       = types.IDX_MAX / 2
	BLOCK_SECTOR_SIZE   types.IDX_T = 4096
	BLOCK_HEAD_SIZE     types.IDX_T = 32
	BLOCK_ALLOC_SIZE    types.IDX_T = 256 * K
	BLOCK_DATA_SIZE                 = BLOCK_ALLOC_SIZE - BLOCK_HEAD_SIZE
	BLOCK_CHECKSUM_SIZE             = BLOCK_HEAD_SIZE
)
