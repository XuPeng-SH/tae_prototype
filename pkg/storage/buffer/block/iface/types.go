package iface

import (
	"io"
	"sync"
	"sync/atomic"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
)

type BlockState = uint8

const (
	BLOCK_LOADED BlockState = iota
	BLOCK_UNLOAD
)

type BlockRTState = uint32

const (
	BLOCK_RT_RUNNING BlockRTState = iota
	BLOCK_RT_CLOSED
)

func AtomicLoadRTState(addr *BlockRTState) BlockRTState {
	return atomic.LoadUint32(addr)
}

func AtomicStoreRTState(addr *BlockRTState, val BlockRTState) {
	atomic.StoreUint32(addr, val)
}

func AtomicCASRTState(addr *BlockRTState, old, new BlockRTState) bool {
	return atomic.CompareAndSwapUint32(addr, old, new)
}

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
	// If the current Refs is already 0, it returns false, else true
	UnRef() bool
	// If the current Refs is not 0, it returns true, else false
	HasRef() bool
}

type IBufferHandle interface {
	io.Closer
	GetID() layout.BlockId
}
