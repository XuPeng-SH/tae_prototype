package vmask

import (
	// "bytes"
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"testing"
)

func TestValidityMask1(t *testing.T) {
	ei := GetEntryIndex(types.IDX_T(32))
	assert.Equal(t, ei.Idx, types.IDX_0)
	assert.Equal(t, ei.Offset, types.IDX_T(32))
	ei = GetEntryIndex(types.IDX_T(82))
	assert.Equal(t, ei.Idx, types.IDX_1)
	assert.Equal(t, ei.Offset, types.IDX_T(18))

	vm1 := New(100)
	assert.Equal(t, vm1.Len(), types.IDX_2)

	var entry EntryT = 0x5
	assert.True(t, entry.IsValid(0))
	assert.False(t, entry.IsValid(1))
	assert.True(t, entry.IsValid(2))
	assert.False(t, entry.IsValid(3))

	entry = vm1.GetEntry(20)
	assert.True(t, entry.AllValid())

	entry = vm1.GetEntry(1)
	assert.Equal(t, vm1.Len(), types.IDX_2)

	row_id := types.IDX_T(50)
	// t.Logf("vm1 row %d valid: %v", row_id, vm1.IsRowValid(row_id))
	assert.True(t, vm1.IsRowValid(row_id))
	vm1.SetInvalid(row_id)
	t.Logf("vm1 row %d valid: %v", row_id, vm1.IsRowValid(row_id))
	assert.False(t, vm1.IsRowValid(row_id))
	e := vm1.GetEntry(GetEntryIndex(row_id).Idx)
	t.Logf("entry: %s", e.String())
	assert.True(t, vm1.IsRowValid(row_id+1))
	vm1.SetValid(row_id)
	assert.True(t, vm1.IsRowValid(row_id))

	vm1.Reset()
	assert.Equal(t, vm1.Len(), types.IDX_0)
}

func TestValidityMask2(t *testing.T) {
	vm1 := New(0)
	assert.Equal(t, vm1.Len(), types.IDX_0)

	vm2 := New(200)
	assert.Equal(t, vm2.Len(), types.IDX_T(4))
	assert.True(t, vm2.IsRowValid(99))
	assert.True(t, vm2.IsRowValid(100))
	assert.True(t, vm2.IsRowValid(127))
	assert.True(t, vm2.IsRowValid(128))
	vm2.InvalidateRows(100)
	assert.False(t, vm2.IsRowValid(99))
	assert.False(t, vm2.IsRowValid(100))
	assert.False(t, vm2.IsRowValid(127))
	assert.True(t, vm2.IsRowValid(128))

	t.Log(vm2.String(129))
	vm2.ValidateRows(100)
	assert.True(t, vm2.IsRowValid(99))
	assert.True(t, vm2.IsRowValid(100))
	assert.True(t, vm2.IsRowValid(127))
	assert.True(t, vm2.IsRowValid(128))
}
