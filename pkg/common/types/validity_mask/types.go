package validity_mask

import (
	"tae/pkg/common/types/constants"
	"unsafe"
)

type EntryT uint64

const (
	BYTES_PER_ENTRY      = (int)(unsafe.Sizeof(EntryT(0)))
	BITS_PER_ENTRY       = BYTES_PER_ENTRY * 8
	MAX_ENTRY            = ^EntryT(0)
	STANDARD_ENTRY_COUNT = (constants.STANDARD_VECTOR_SIZE + (BITS_PER_ENTRY - 1)) / BITS_PER_ENTRY
)

type EntryIndex struct {
	Idx    int
	Offset int
}

type ValidityMask struct {
	Data []EntryT
}
