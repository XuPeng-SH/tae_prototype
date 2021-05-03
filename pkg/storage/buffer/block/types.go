package block

import (
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
)

type IBlockBuffer interface {
	buf.IBuffer
	GetID() types.IDX_T
}

type BlockBuffer struct {
	buf.IBuffer
	ID types.IDX_T
}
