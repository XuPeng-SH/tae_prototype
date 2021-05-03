package buffer

import (
	"hash"
	"io"
	"tae/pkg/common/types"
)

type CheckSumT uint64

type IBuffer interface {
	ReadAt(r io.ReaderAt, off int64) (n int, err error)
	WriteAt(w io.WriterAt, off int64) (n int, err error)
	// Clear()
	// Capacity() types.IDX_T
}

type Buffer struct {
	Data       []byte
	DataSize   types.IDX_T
	HeaderSize types.IDX_T
	Hasher     hash.Hash
}
