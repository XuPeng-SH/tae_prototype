package vector_buffer

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"testing"
)

func TestInteger(t *testing.T) {
	vf := NewVectorBuffer(WithItemType(types.LT_INTEGER), WithSize((int)(types.LT_INTEGER.GetPhysicalType().Size())))
	assert.Equal(t, vf.Size(), (int)(types.PSIZE_INT32))

	v0 := int32(-33)
	vf.SetValue(0, v0)
	v1 := vf.GetValue(0)
	assert.Equal(t, v1, v0)

	vf2 := NewVectorBuffer(WithItemType(types.LT_UINTEGER), WithSize((int)(types.LT_UINTEGER.GetPhysicalType().Size())))
	assert.Equal(t, vf2.Size(), (int)(types.PSIZE_UINT32))

	v2 := uint32(23)
	vf2.SetValue(0, v2)
	v3 := vf2.GetValue(0)
	assert.Equal(t, v3, v2)
}

func TestBoolean(t *testing.T) {
	vf1 := NewVectorBuffer(WithItemType(types.LT_BOOLEAN), WithSize((int)(types.LT_BOOLEAN.GetPhysicalType().Size())))
	assert.Equal(t, vf1.Size(), (int)(types.PSIZE_BOOL))
	v0 := true
	vf1.SetValue(0, v0)
	v1 := vf1.GetValue(0)
	assert.Equal(t, v1, v0)
}

func TestFloat32(t *testing.T) {
	vf1 := NewVectorBuffer(WithItemType(types.LT_FLOAT32), WithSize((int)(types.LT_FLOAT32.GetPhysicalType().Size())))
	assert.Equal(t, vf1.Size(), (int)(types.PSIZE_FLOAT32))
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
