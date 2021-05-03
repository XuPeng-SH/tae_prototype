package layout

import (
	"fmt"
	"tae/pkg/common/types"
)

const (
	K                            types.IDX_T = 1024
	M                            types.IDX_T = 1024 * K
	G                            types.IDX_T = 1024 * M
	MAIN_HEADER_MAGIC                        = "BASE"
	MAIN_HEADER_BYTES_MAGIC                  = 4
	MAIN_HEADER_BYTES_RESERVERED             = 28
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

var (
	INVALID_BLOCK_ID        = BlockId{Part: ^uint32(0), Offset: ^uint32(0)}
	MIN_TRANSIENT_BLOCK_ID  = BlockId{Part: ^uint32(0) / 2, Offset: uint32(0)}
	MAX_PERSISTENT_BLOCK_ID = BlockId{Part: MIN_TRANSIENT_BLOCK_ID.Part - 1, Offset: ^uint32(0)}
)

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
