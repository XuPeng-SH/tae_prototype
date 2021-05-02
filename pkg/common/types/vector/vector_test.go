package vector

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	"tae/pkg/common/types/value"
	vbuff "tae/pkg/common/types/vbuff"
	"testing"
)

func TestInitWithLogicType(t *testing.T) {
	vec := NewVector(WithInitByLogicType(types.LT_FLOAT32))
	t.Logf("vec itemsize=%d", vec.GetBuffer().GetItemSize())
	t.Logf("vec itemcount=%d", vec.GetBuffer().MaxItems())
	t.Logf("vec size=%d", vec.GetBuffer().Size())
	t.Log(vec)
	assert.Equal(t, vec.GetBuffer().GetItemSize(), types.PSIZE_FLOAT32)
	assert.Equal(t, vec.GetBuffer().MaxItems(), constants.STANDARD_VECTOR_SIZE)
	assert.Equal(t, vec.GetBuffer().Size(), types.PSIZE_FLOAT32*constants.STANDARD_VECTOR_SIZE)
	assert.Equal(t, vec.GetBuffer().GetType(), vbuff.STANDARD_BUFFER)
	assert.Equal(t, vec.GetType(), FLAT_VECTOR)
}

func TestInitWithValue(t *testing.T) {
	f_val := float32(-1.567)
	val := value.NewValue(f_val)
	vec := NewVector(WithInitByValue(val))
	t.Logf("vec itemsize=%d", vec.GetBuffer().GetItemSize())
	t.Logf("vec itemcount=%d", vec.GetBuffer().MaxItems())
	t.Logf("vec size=%d", vec.GetBuffer().Size())
	t.Log(vec)
	assert.Equal(t, vec.GetBuffer().GetItemSize(), types.PSIZE_FLOAT32)
	assert.Equal(t, vec.GetBuffer().MaxItems(), types.IDX_1)
	assert.Equal(t, vec.GetBuffer().Size(), types.PSIZE_FLOAT32)
	assert.Equal(t, vec.GetBuffer().GetType(), vbuff.STANDARD_BUFFER)
	assert.Equal(t, vec.GetType(), CONSTANT_VECTOR)
}
