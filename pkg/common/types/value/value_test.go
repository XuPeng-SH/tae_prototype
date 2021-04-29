package value

import (
	"testing"
)

func TestAll(t *testing.T) {
	v := NewValue(-22)
	t.Log(v.Val.Data)
	t.Log(v.Type.LType)

	v2 := NewValue(int16(32))
	t.Log(v2.Val.Data)
	t.Log(v2.Type.LType)
}
