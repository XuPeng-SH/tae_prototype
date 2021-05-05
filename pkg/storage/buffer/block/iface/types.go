package iface

import (
	"io"
	"sync"
	"sync/atomic"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
)

type BlockState = uint32

const (
	BLOCK_UNLOAD BlockState = iota
	BLOCK_LOADING
	BLOCK_ROOLBACK
	BLOCK_COMMIT
	BLOCK_LOADED
)

func AtomicLoadState(addr *BlockState) BlockState {
	return atomic.LoadUint32(addr)
}

func AtomicStoreState(addr *BlockState, val BlockState) {
	atomic.StoreUint32(addr, val)
}

func AtomicCASState(addr *BlockState, old, new BlockState) bool {
	return atomic.CompareAndSwapUint32(addr, old, new)
}

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
	PrepareLoad() bool
	RollbackLoad()
	CommitLoad() error
	MakeHandle() IBufferHandle
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
	SetBuffer(buffer buf.IBuffer) error
}

type IBufferHandle interface {
	io.Closer
	GetID() layout.BlockId
}
