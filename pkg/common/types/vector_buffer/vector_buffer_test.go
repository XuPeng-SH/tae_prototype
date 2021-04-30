package vector_buffer

import (
	"github.com/stretchr/testify/assert"
	"tae/pkg/common/types"
	"testing"
)

func TestAll(t *testing.T) {
	vf := NewVectorBuffer(WithItemType(types.LT_INTEGER), WithSize((int)(types.LT_INTEGER.GetPhysicalType().Size())))
	assert.Equal(t, vf.Size(), (int)(types.PSIZE_INT32))
	t.Logf("vf.size()=%d", vf.Size())

	vf.SetValue(0, int32(33))
}
