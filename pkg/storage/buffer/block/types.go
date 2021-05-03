package block

import (
	"sync"
	"tae/pkg/common/types"
	buf "tae/pkg/storage/buffer"
	"tae/pkg/storage/layout"
)

type IBlockBuffer interface {
	buf.IBuffer
	GetID() types.IDX_T
}

type BlockBuffer struct {
	buf.IBuffer
	ID types.IDX_T
}

type BlockState uint8

const (
	BLOCK_LOADED BlockState = iota
	BLOCK_UNLOAD
)

type BlockHandleCtx struct {
	ID          layout.BlockId
	Buff        buf.IBuffer
	Destroyable bool
}

type IBlockHandle interface {
	sync.Locker
	// GetID() layout.BlockId
	// Unload()
	// Loadable() bool
	// GetBuff() buf.IBuffer
	// GetState() BlockState
	// Size() types.IDX_T
	// IsDestroyable() bool
}

type BlockHandle struct {
	sync.Mutex
	State       BlockState
	ID          layout.BlockId
	Buff        buf.IBuffer
	Destroyable bool
	Capacity    types.IDX_T
}
