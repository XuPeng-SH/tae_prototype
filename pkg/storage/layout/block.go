package layout

import (
	// "fmt"
	"tae/pkg/common/types"
)

const (
	BLOCK_INVALID                    = types.IDX_MAX
	BLOCK_PERSISTENT_MAX             = types.IDX_MAX / 2
	BLOCK_TRANSIENT_MIN              = BLOCK_PERSISTENT_MAX
	BLOCK_SECTOR_SIZE    types.IDX_T = 4 * K
	BLOCK_HEAD_SIZE      types.IDX_T = 32
	BLOCK_ALLOC_SIZE     types.IDX_T = 256 * K
	BLOCK_DATA_SIZE                  = BLOCK_ALLOC_SIZE - BLOCK_HEAD_SIZE
	BLOCK_CHECKSUM_SIZE              = BLOCK_HEAD_SIZE
)
