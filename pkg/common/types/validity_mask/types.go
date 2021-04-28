package validity_mask

import (
	"unsafe"
)

type EntryT uint64

const (
	BYTES_PER_ENTRY = (int)(unsafe.Sizeof(EntryT(0)))
	BITS_PER_ENTRY  = BYTES_PER_ENTRY * 8
	MAX_ENTRY       = ^EntryT(0)
)

type EntryIndex struct {
	Idx    int
	Offset int
}

type ByteIndex struct {
	Idx    int
	Offset int
}

type ValidityMask struct {
	Data []byte
}
