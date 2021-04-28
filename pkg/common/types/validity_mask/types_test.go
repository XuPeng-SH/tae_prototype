package validity_mask

import (
	"testing"
)

func TestType(t *testing.T) {
	t.Log(MAX_ENTRY)
	t.Log(BITS_PER_ENTRY)
	t.Log(WhichEntry(20))
	t.Log(WhichEntry(80))
	t.Log(WhichEntry(180))
}
