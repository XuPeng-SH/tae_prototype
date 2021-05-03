package buffer

import (
	"fmt"
	"hash"
	"io"
	"tae/pkg/common/types"
)

type BufferType uint8

const (
	NA_BUFFER BufferType = iota
	BLOCK_BUFFER
)

func (bt BufferType) String() string {
	switch bt {
	case BLOCK_BUFFER:
		return "BLOCK_BUFFER"
	}
	panic(fmt.Sprintf("UNKNOWN buffer type %d", bt))
}

type IBuffer interface {
	ReadAt(r io.ReaderAt, off int64) (n int, err error)
	WriteAt(w io.WriterAt, off int64) (n int, err error)
	Clear()
	Capacity() int64
	GetType() BufferType
	// String() string
	// ToString(opts...interface{}) string
}

type Buffer struct {
	Data       []byte
	DataSize   types.IDX_T
	HeaderSize types.IDX_T
	Hasher     hash.Hash
	Type       BufferType
}
