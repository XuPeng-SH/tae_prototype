package selection_vector

import (
	"testing"
)

func TestInit(t *testing.T) {
	op := WithSize(64)
	sv := New(op)
	t.Log(sv.Data.Cap())
	t.Log(sv.Data.Len())
}
