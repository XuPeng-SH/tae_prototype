package manager

import (
	// "sync/atomic"
	"sync"
	"tae/pkg/common/types"
	// buf "tae/pkg/storage/buffer"
	blk "tae/pkg/storage/buffer/block"
	"tae/pkg/storage/layout"
)

type IBufferManager interface {
	sync.Locker

	// GetSize() types.IDX_T
	// GetCapacity() types.IDX_T
	// SetCapacity(c types.IDX_T)

	RegisterBlock(blk_id layout.BlockId) blk.IBlockHandle
	// UnregisterBlock(blk_id layout.BlockId, can_destroy bool)

	// RegisterMemory(blk_id layout.BlockId, can_destroy bool) blk.IBlockHandle
	// // Allocate(size types.IDX_T) buf.IBufferH

	// Pin(h blk.IBlockHandle) buf.IBuffer
	// Unpin(h blk.IBlockHandle)
}

type BufferManager struct {
	sync.Mutex
	Size        types.IDX_T
	Capacity    types.IDX_T
	Blocks      map[layout.BlockId]blk.IBlockHandle // Manager is not responsible to Close handle
	TransientID layout.BlockId

	// TempPath string
}
