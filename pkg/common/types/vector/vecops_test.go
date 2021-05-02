package vector

import (
	"github.com/stretchr/testify/assert"
	// "tae/pkg/common/types"
	// "tae/pkg/common/types/constants"
	"tae/pkg/common/types"
	"tae/pkg/common/types/value"
	// vbuff "tae/pkg/common/types/vbuff"
	"testing"
)

func TestCopyConstVector(t *testing.T) {
	val1 := value.NewValue(-1.222)
	// val2 := value.NewValue(2.3)
	src := NewVector(WithInitByValue(val1))
	des := NewVector()
	panic1 := func() {
		Copy(src, des, 0, 0, val1.GetPhysicalTypeSize())
	}
	assert.Panics(t, panic1)

	des = NewVector(WithInitByLogicType(types.LT_BIGINT))
	panic1 = func() {
		Copy(src, des, 0, 0, val1.GetPhysicalTypeSize())
	}
	assert.Panics(t, panic1)

	des = NewVector(WithInitByLogicType(val1.GetLogicType()))
	t.Log(src)
	t.Log(des)
	Copy(src, des, 0, 0, val1.GetPhysicalTypeSize())
}
