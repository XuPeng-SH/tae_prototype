package block

import (
	"sync"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	blkif "tae/pkg/storage/buffer/block/iface"
	mgrif "tae/pkg/storage/buffer/manager/iface"
	"tae/pkg/storage/layout"
)

type BlockBuffer struct {
	buf.IBuffer
	ID layout.BlockId
}

type BlockHandleCtx struct {
	ID          layout.BlockId
	Buff        buf.IBuffer
	Destroyable bool
	Manager     mgrif.IBufferManager
}

type BlockHandle struct {
	sync.Mutex
	State       blkif.BlockState
	ID          layout.BlockId
	Buff        buf.IBuffer
	Destroyable bool
	Capacity    types.IDX_T
	RTState     blkif.BlockRTState
	Refs        types.IDX_T
	Manager     mgrif.IBufferManager
}

// BufferHandle is created from IBufferManager::Pin, which will set the IBlockHandle reference to 1
// The following IBufferManager::Pin will call IBlockHandle::Ref to increment the reference count
// BufferHandle should alway be closed manually when it is not needed, which will call IBufferManager::Unpin
type BufferHandle struct {
	Handle  blkif.IBlockHandle
	Manager mgrif.IBufferManager
}
