package layout

import (
	"fmt"
	"tae/pkg/common/types"
)

const (
	BLOCK_INVALID = types.IDX_MAX
	BLOCK_MAX     = types.IDX_MAX / 2
)

const (
	MAIN_HEADER_MAGIC            = "BASE"
	MAIN_HEADER_BYTES_MAGIC      = 4
	MAIN_HEADER_BYTES_RESERVERED = 28
)

func init() {
	if len(MAIN_HEADER_MAGIC) != MAIN_HEADER_BYTES_MAGIC {
		panic(fmt.Sprintf("Main header magic size should be %d", MAIN_HEADER_BYTES_MAGIC))
	}
}

type BlockId struct {
	Part   uint32
	Offset uint32
}

type MainHeader struct {
	Version    types.IDX_T
	Reservered [MAIN_HEADER_BYTES_RESERVERED]byte
}

type DataHeader struct {
	SequenceNumber types.IDX_T
	FirstMetaBlk   BlockId
	FirstFreeBlk   BlockId
	BlockCount     types.IDX_T
	PartCount      types.IDX_T
}
