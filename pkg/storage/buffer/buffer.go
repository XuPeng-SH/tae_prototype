package buffer

import (
	"bytes"
	"crypto/sha256"
	"io"
	"tae/pkg/common/types"
	"tae/pkg/common/util/hack"
	"tae/pkg/storage/layout"
)

var (
	_ IBuffer = (*Buffer)(nil)
)

func NewBuffer(bufsize types.IDX_T, pool IMemoryPool) IBuffer {
	if bufsize%layout.BLOCK_SECTOR_SIZE != 0 {
		bufsize += layout.BLOCK_SECTOR_SIZE - (bufsize % layout.BLOCK_SECTOR_SIZE)
	}
	node := pool.Get(bufsize)
	buf := &Buffer{
		Node:       node,
		HeaderSize: layout.BLOCK_HEAD_SIZE,
		DataSize:   bufsize - layout.BLOCK_HEAD_SIZE,
		Hasher:     sha256.New(),
	}
	return buf
}

func (buf *Buffer) ReadAt(r io.ReaderAt, off int64) (n int, err error) {
	if n, err = r.ReadAt(buf.Node.Data, off); err != nil {
		return n, err
	}
	buf.Hasher.Reset()
	if n, err = buf.Hasher.Write(buf.Node.Data[buf.HeaderSize:]); err != nil {
		return n, err
	}
	if !bytes.Equal(buf.Hasher.Sum(nil), buf.Node.Data[:buf.HeaderSize]) {
		panic("CheckSum mismatch")
	}
	return n, err
}

func (buf *Buffer) WriteAt(w io.WriterAt, off int64) (n int, err error) {
	buf.Hasher.Reset()
	if n, err = buf.Hasher.Write(buf.Node.Data[buf.HeaderSize:]); err != nil {
		return n, err
	}
	copy(buf.Node.Data, buf.Hasher.Sum(nil))
	n, err = w.WriteAt(buf.Node.Data, off)
	return n, err
}

func (buf *Buffer) GetType() BufferType {
	return buf.Type
}

func (buf *Buffer) Close() error {
	buf.Node.Pool.Put(buf.Node)
	return nil
}

func (buf *Buffer) Clear() {
	hack.MemsetRepeatByte(buf.Node.Data, byte(0))
}

func (buf *Buffer) Capacity() int64 {
	return int64(buf.DataSize + buf.HeaderSize)
}
