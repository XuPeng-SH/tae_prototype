package vbuff

import (
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	// "tae/pkg/common/types/value"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInteger(t *testing.T) {
	vf := NewVectorBuffer(WithItemType(types.LT_INTEGER), WithSize((types.IDX_T)(types.LT_INTEGER.GetPhysicalType().Size())))
	assert.Equal(t, vf.Size(), types.PSIZE_INT32)

	v0 := int32(-33)
	vf.SetValue(types.IDX_0, v0)
	v1 := vf.GetValue(types.IDX_0)
	assert.Equal(t, v1, v0)

	vf2 := NewVectorBuffer(WithItemType(types.LT_UINTEGER), WithSize(types.LT_UINTEGER.GetPhysicalType().Size()))
	assert.Equal(t, vf2.Size(), types.PSIZE_UINT32)

	v2 := uint32(23)
	vf2.SetValue(0, v2)
	v3 := vf2.GetValue(0)
	assert.Equal(t, v3, v2)
}

func TestBoolean(t *testing.T) {
	vf1 := NewVectorBuffer(WithItemType(types.LT_BOOLEAN), WithSize(types.LT_BOOLEAN.GetPhysicalType().Size()))
	assert.Equal(t, vf1.Size(), types.PSIZE_BOOL)
	v0 := true
	vf1.SetValue(0, v0)
	v1 := vf1.GetValue(0)
	assert.Equal(t, v1, v0)
}

func TestFloat32(t *testing.T) {
	vf1 := NewVectorBuffer(WithItemType(types.LT_FLOAT32), WithSize(types.LT_FLOAT32.GetPhysicalType().Size()))
	assert.Equal(t, vf1.Size(), types.PSIZE_FLOAT32)
	v0 := float32(-120.34)
	vf1.SetValue(0, v0)
	v1 := vf1.GetValue(0)
	assert.Equal(t, v1, v0)
}

func TestDictBuff(t *testing.T) {
	dbuff := NewDictonaryBuffer(WithDictBuffCount(20), WithDictBuffItemType(types.LT_FLOAT32))
	assert.Equal(t, dbuff.GetType(), DICTIONARY_BUFFER)
	assert.Equal(t, dbuff.GetItemType(), types.LT_FLOAT32)
	t.Log(dbuff.SelVec.ToString(10))
}

func TestForceRepeat(t *testing.T) {
	vf := NewVectorBuffer(WithItemType(types.LT_FLOAT64),
		WithSize(types.LT_FLOAT64.GetPhysicalType().Size()*constants.STANDARD_VECTOR_SIZE))
	fval := float64(0.232)
	idx := types.IDX_T(2)
	vf.SetValue(idx, fval)
	assert.Equal(t, vf.GetValue(idx), fval)
	t.Log(vf)
	count := types.IDX_T(5)
	vf.ForceRepeat(idx, count)
	t.Log(vf)
	for i := idx; i < idx+count; i++ {
		assert.Equal(t, vf.GetValue(i), fval)
	}
	assert.Equal(t, vf.GetValue(idx+count), float64(0))
}
