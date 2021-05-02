package selvec

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"testing"
)

func TestAll(t *testing.T) {
	count := types.IDX_T(16)
	op := WithCount(count)
	sv := New(op)
	t.Log(sv.ToString(count))
	assert.Equal(t, sv.Count(), count)
	sv.SetIndex(0, 2)
	sv.SetIndex(1, 3)
	sv.SetIndex(2, 3)
	sv.SetIndex(3, 4)
	t.Log(sv.ToString(8))
	assert.Equal(t, (int)(sv.GetIndex(0)), 2)
	assert.Equal(t, (int)(sv.GetIndex(1)), 3)
	assert.Equal(t, (int)(sv.GetIndex(2)), 3)
	assert.Equal(t, (int)(sv.GetIndex(3)), 4)

	sv.Swap(0, 3)
	t.Log(sv.ToString(8))
	assert.Equal(t, (int)(sv.GetIndex(0)), 4)
	assert.Equal(t, (int)(sv.GetIndex(3)), 2)

	sv2 := New()
	assert.True(t, sv2.Empty())
	sv2.InitWithCount(4)
	assert.Equal(t, sv2.Count(), types.IDX_T(4))

	sv2.SetIndex(0, 0)
	sv2.SetIndex(1, 2)
	sv2.SetIndex(2, 4)
	sv2.SetIndex(3, 6)
	t.Log(sv2.ToString(4))

	data := sv.Slice(*sv2, 4)
	new_sv := SelectionVector{
		Data: data,
	}
	t.Log(new_sv.ToString(new_sv.Count()))
	assert.Equal(t, new_sv.GetIndex(0), EntryT(4))
	assert.Equal(t, new_sv.GetIndex(1), EntryT(3))

	t.Log(ZERO_SV.ToString(10))
	t.Log(SEQUENTIAL_SV.ToString(10))
}
