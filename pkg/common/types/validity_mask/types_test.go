package validity_mask

import (
	// "bytes"
	// "encoding/binary"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidityMask1(t *testing.T) {
	ei := GetEntryIndex(32)
	assert.Equal(t, ei.Idx, 0)
	assert.Equal(t, ei.Offset, 32)
	ei = GetEntryIndex(82)
	assert.Equal(t, ei.Idx, 1)
	assert.Equal(t, ei.Offset, 18)

	vm1 := New(100)
	assert.Equal(t, vm1.Len(), 2)

	var entry EntryT = 0x5
	assert.True(t, entry.IsValid(0))
	assert.False(t, entry.IsValid(1))
	assert.True(t, entry.IsValid(2))
	assert.False(t, entry.IsValid(3))

	entry = vm1.GetEntry(20)
	assert.True(t, entry.AllValid())

	entry = vm1.GetEntry(1)
	assert.Equal(t, vm1.Len(), 2)

	row_id := 50
	// t.Logf("vm1 row %d valid: %v", row_id, vm1.IsRowValid(row_id))
	assert.True(t, vm1.IsRowValid(row_id))
	vm1.SetInvalid(row_id)
	// t.Logf("vm1 row %d valid: %v", row_id, vm1.IsRowValid(row_id))
	assert.False(t, vm1.IsRowValid(row_id))
	e := vm1.GetEntry(GetEntryIndex(row_id).Idx)
	t.Logf("entry: %v", e.ToString())
	assert.True(t, vm1.IsRowValid(row_id+1))
	vm1.SetValid(row_id)
	assert.True(t, vm1.IsRowValid(row_id))

	vm1.Reset()
	assert.Equal(t, vm1.Len(), 0)

	// buff := new(bytes.Buffer)
	// buff.WriteString("hello")
	// t.Log(buff.Cap())
	// t.Log(buff.Len())
	// t.Log(buff.String())
	// buff2 := new(bytes.Buffer)
	// var num uint16 = 12
	// if err := binary.Write(buff2, binary.LittleEndian, num); err != nil {
	// 	t.Error(err)
	// }
	// t.Log(buff2.Cap())
	// t.Log(buff2.Len())
	// t.Log(buff2.String())
}

func TestValidityMask2(t *testing.T) {
	vm1 := New(0)
	assert.Equal(t, vm1.Len(), 0)

	vm2 := New(200)
	assert.Equal(t, vm2.Len(), 4)
	assert.True(t, vm2.IsRowValid(99))
	assert.True(t, vm2.IsRowValid(100))
	assert.True(t, vm2.IsRowValid(127))
	assert.True(t, vm2.IsRowValid(128))
	vm2.InvalidateRows(100)
	assert.False(t, vm2.IsRowValid(99))
	assert.False(t, vm2.IsRowValid(100))
	assert.False(t, vm2.IsRowValid(127))
	assert.True(t, vm2.IsRowValid(128))
	t.Log(vm2.ToString(129))
	vm2.ValidateRows(100)
	assert.True(t, vm2.IsRowValid(99))
	assert.True(t, vm2.IsRowValid(100))
	assert.True(t, vm2.IsRowValid(127))
	assert.True(t, vm2.IsRowValid(128))
}
