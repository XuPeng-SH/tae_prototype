package iface

import (
	"sync"
	"tae/pkg/common/types"
	blkif "tae/pkg/storage/buffer/block/iface"
	"tae/pkg/storage/layout"
)

type IBufferManager interface {
	sync.Locker

	GetUsageSize() types.IDX_T
	GetCapacity() types.IDX_T
	SetCapacity(c types.IDX_T)

	RegisterBlock(blk_id layout.BlockId) blkif.IBlockHandle
	UnregisterBlock(blk_id layout.BlockId, can_destroy bool)

	// RegisterMemory(blk_id layout.BlockId, can_destroy bool) blk.IBlockHandle
	// // Allocate(size types.IDX_T) buf.IBufferH

	Pin(h blkif.IBlockHandle) blkif.IBufferHandle
	Unpin(h blkif.IBlockHandle)
}
