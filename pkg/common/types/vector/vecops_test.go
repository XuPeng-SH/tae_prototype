package vector

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"tae/pkg/common/types/value"
	"testing"
)

func TestCopyConstVector(t *testing.T) {
	val1 := value.NewValue(-1.222)
	src := NewVector(WithInitByValue(val1))
	des := NewVector()
	panic1 := func() {
		Copy(src, des, 0, 0, 1)
	}
	assert.Panics(t, panic1)

	des = NewVector(WithInitByLogicType(types.LT_BIGINT))
	panic1 = func() {
		Copy(src, des, 0, 0, 1)
	}
	assert.Panics(t, panic1)

	des = NewVector(WithInitByLogicType(val1.GetLogicType()))
	count := types.IDX_T(3)
	Copy(src, des, 0, 0, count)
	for i := types.IDX_0; i < count; i++ {
		assert.Equal(t, des.GetValue(i), val1.GetValue())
	}
	assert.NotEqual(t, des.GetValue(count), val1.GetValue())
}

func TestCopyFlatVector(t *testing.T) {
	src := NewVector(WithInitByLogicType(types.LT_FLOAT32))
	for i := types.IDX_0; i < src.GetBuffer().MaxItems(); i++ {
		val := value.NewValue(float32(i + 10000))
		src.SetValue(i, val)
	}
	des := NewVector(WithInitByLogicType(src.GetLogicType()))
	count := src.GetBuffer().MaxItems()
	Copy(src, des, 0, 0, count)
	assert.Equal(t, des.GetValue(types.IDX_T(1000)), float32(1000+10000))

	des = NewVector(WithInitByLogicType(src.GetLogicType()))
	Copy(src, des, 0, 1000, count-1000)
	assert.Equal(t, des.GetValue(types.IDX_T(1000)), float32(10000))

	des = NewVector(WithInitByLogicType(src.GetLogicType()))
	panic1 := func() {
		Copy(src, des, 0, count-10, types.IDX_T(11))
	}
	assert.Panics(t, panic1)
	Copy(src, des, 0, count/2, types.IDX_T(10))
	Copy(src, des, 0, count/2+30, types.IDX_T(10))
	assert.Equal(t, des.GetValue(count/2), float32(10000))
	assert.Equal(t, des.GetValue(count/2+9), float32(10009))
	assert.Equal(t, des.GetValue(count/2+10), float32(0))
	assert.Equal(t, des.GetValue(count/2+30), float32(10000))
	assert.Equal(t, des.GetValue(count/2+39), float32(10009))
	assert.Equal(t, des.GetValue(count/2+40), float32(0))
}
