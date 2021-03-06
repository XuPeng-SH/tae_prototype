package value

import (
	"testing"
)

func TestAll(t *testing.T) {
	v := NewValue(-22)
	t.Log(v.Data)
	t.Log(v.Type.LType)
	t.Log(v.String())

	v2 := NewValue(int16(32))
	t.Log(v2.Data)
	t.Log(v2.Type.LType)
	t.Log(v2.String())

	v3 := NewValue()
	t.Log(v3.Data)
	t.Log(v3.Type.LType)
	t.Log(v3.String())

	v4 := NewValue(true)
	t.Log(v4.Data)
	t.Log(v4.Type.LType)
	t.Log(v4.String())
}
