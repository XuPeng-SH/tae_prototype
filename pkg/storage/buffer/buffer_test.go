package buffer

import (
	"os"
	"tae/pkg/common/types"
	"tae/pkg/storage/layout"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuffer(t *testing.T) {
	pool := NewSimpleMemoryPool(layout.BLOCK_ALLOC_SIZE * 100)
	node := pool.Get(layout.BLOCK_SECTOR_SIZE, false)
	buf := NewBuffer(node)
	for i := layout.BLOCK_HEAD_SIZE; i < layout.BLOCK_SECTOR_SIZE; i++ {
		buf.(*Buffer).Node.Data[i] = byte((i - layout.BLOCK_HEAD_SIZE) % 256)
	}
	assert.Equal(t, NA_BUFFER, buf.GetType())

	path := "/tmp/tttttt"
	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	assert.Equal(t, err, nil)
	_, err = buf.WriteAt(w, 0)
	assert.Equal(t, err, nil)
	w.Close()
	// t.Log(buf.Data)
	r, err := os.OpenFile(path, os.O_RDONLY, 0666)
	assert.Equal(t, err, nil)
	node2 := pool.Get(layout.BLOCK_SECTOR_SIZE, false)
	buf2 := NewBuffer(node2)
	_, err = buf2.ReadAt(r, 0)
	assert.Equal(t, err, nil)
	r.Close()
	// t.Log(buf2.Data)

	assert.Equal(t, buf.Capacity(), int64(layout.BLOCK_SECTOR_SIZE))
	buf2.Clear()
	assert.Equal(t, buf2.(*Buffer).Node.Data[22], byte(0))
	assert.Equal(t, buf2.(*Buffer).Node.Data[23], byte(0))
	assert.Equal(t, buf2.(*Buffer).Node.Data[24], byte(0))
}

func TestBufferPool(t *testing.T) {
	pool_size := layout.BLOCK_ALLOC_SIZE
	pool := NewSimpleMemoryPool(pool_size)
	node := pool.Get(types.IDX_T(100), false)
	buf := NewBuffer(node)
	buf_cap := types.IDX_T(buf.Capacity())
	assert.Equal(t, pool.GetUsageSize(), buf_cap)
	t.Log(pool.GetUsageSize())
	t.Log(pool.GetCapacity())
	buf.Close()
	t.Log(pool.GetUsageSize())
	t.Log(pool.GetCapacity())
	assert.Equal(t, pool.GetUsageSize(), types.IDX_0)
	assert.Equal(t, buf.Capacity(), int64(0))

	node2 := pool.Get(pool_size+1, false)
	buf2 := NewBuffer(node2)
	assert.Equal(t, buf2, nil)
}
