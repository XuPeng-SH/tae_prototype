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

func TestFlatVector(t *testing.T) {
	src := NewVector(WithInitByLogicType(types.LT_FLOAT32))
	for i := types.IDX_0; i < 4; i++ {
		val := value.NewValue(float32(i))
		src.SetValue(i, val)
		val2 := src.GetValue(i)
		assert.Equal(t, val.GetValue(), val2)
	}
	t.Log(src)
}

func TestNormality(t *testing.T) {
	fval := float32(1.238)
	val := value.NewValue(fval)
	const_vec := NewVector(WithInitByValue(val))
	t.Log(const_vec)
	assert.Equal(t, const_vec.GetType(), CONSTANT_VECTOR)
	count := types.IDX_T(10)
	const_vec.Flatten(count)
	assert.Equal(t, const_vec.GetType(), FLAT_VECTOR)
	t.Log(const_vec)
	for i := types.IDX_0; i < count; i++ {
		assert.Equal(t, const_vec.GetValue(i), fval)
	}
	assert.Equal(t, const_vec.GetValue(count), float32(0))
}

func TestSequenceVector(t *testing.T) {
	src := NewVector(WithInitByLogicType(types.LT_FLOAT32))
	panic1 := func() {
		src.GetSequence()
	}
	assert.Panics(t, panic1)
	assert.Equal(t, src.GetType(), FLAT_VECTOR)
	seq := SequenceData{
		Start: 0,
		Step:  1,
	}
	src.ToSeqenceVector(&seq)
	t.Log(seq.String())
	seq2 := src.GetSequence()
	assert.Equal(t, seq.Start, seq2.Start)
	assert.Equal(t, seq.Step, seq2.Step)
}
