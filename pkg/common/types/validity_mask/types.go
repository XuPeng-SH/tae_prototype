package validity_mask

import (
	"bytes"
	"unsafe"
)

type EntryT uint64

const (
	BITS_PER_ENTRY = (uint)(unsafe.Sizeof(EntryT(0))) * 8
	MAX_ENTRY      = ^EntryT(0)
)

type ValidityMask struct {
	Data *bytes.Buffer
}
