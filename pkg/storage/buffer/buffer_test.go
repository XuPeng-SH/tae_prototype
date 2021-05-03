package buffer

import (
	"github.com/stretchr/testify/assert"
	"os"
	"tae/pkg/storage/layout"
	"testing"
)

func TestBuffer(t *testing.T) {
	buf := NewBuffer(layout.BLOCK_SECTOR_SIZE)
	for i := layout.BLOCK_HEAD_SIZE; i < layout.BLOCK_SECTOR_SIZE; i++ {
		buf.(*Buffer).Data[i] = byte((i - layout.BLOCK_HEAD_SIZE) % 256)
	}

	path := "/tmp/tttttt"
	w, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	assert.Equal(t, err, nil)
	_, err = buf.WriteAt(w, 0)
	assert.Equal(t, err, nil)
	w.Close()
	// t.Log(buf.Data)
	r, err := os.OpenFile(path, os.O_RDONLY, 0666)
	assert.Equal(t, err, nil)
	buf2 := NewBuffer(layout.BLOCK_SECTOR_SIZE)
	_, err = buf2.ReadAt(r, 0)
	assert.Equal(t, err, nil)
	r.Close()
	// t.Log(buf2.Data)
}
