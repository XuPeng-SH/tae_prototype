package block

import (
	"io"
	"sync"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
)

type IBlockBuffer interface {
	buf.IBuffer
	GetID() layout.BlockId
}

type BlockBuffer struct {
	buf.IBuffer
	ID layout.BlockId
}

type BlockState uint8

const (
	BLOCK_LOADED BlockState = iota
	BLOCK_UNLOAD
)

type BlockRTState uint8

const (
	BLOCK_RT_RUNNING BlockRTState = iota
	BLOCK_RT_CLOSED
)

type BlockHandleCtx struct {
	ID          layout.BlockId
	Buff        buf.IBuffer
	Destroyable bool
}

type IBlockHandle interface {
	sync.Locker
	io.Closer
	GetID() layout.BlockId
	// Unload()
	// Loadable() bool
	// GetBuff() buf.IBuffer
	GetState() BlockState
	GetCapacity() types.IDX_T
	// Size() types.IDX_T
	// IsDestroyable() bool
	IsClosed() bool
	Ref()
	UnRef()
	HasRef() bool
}

type BlockHandle struct {
	sync.Mutex
	State       BlockState
	ID          layout.BlockId
	Buff        buf.IBuffer
	Destroyable bool
	Capacity    types.IDX_T
	RTState     BlockRTState
	Refs        types.IDX_T
}

type IBufferHandle interface {
	io.Closer
}

type BufferHandle struct {
	Handle IBlockBuffer
	Buff   buf.IBuffer
}
