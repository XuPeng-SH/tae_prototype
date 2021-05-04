package iface

import (
	"io"
	"sync"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
)

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

type IBlockBuffer interface {
	buf.IBuffer
	GetID() layout.BlockId
}

type IBlockHandle interface {
	sync.Locker
	io.Closer
	GetID() layout.BlockId
	// Unload()
	// Loadable() bool
	// GetBuff() buf.IBuffer
	Load() IBufferHandle
	GetState() BlockState
	GetCapacity() types.IDX_T
	// Size() types.IDX_T
	// IsDestroyable() bool
	IsClosed() bool
	Ref()
	UnRef()
	HasRef() bool
}

type IBufferHandle interface {
	io.Closer
}
