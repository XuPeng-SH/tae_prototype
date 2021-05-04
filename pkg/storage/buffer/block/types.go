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
}

type BufferHandle struct {
	Handle  blkif.IBlockHandle
	Manager mgrif.IBufferManager
}
