package vmask

import (
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	"unsafe"
)

type EntryT uint64

const (
	BYTES_PER_ENTRY      = (types.IDX_T)(unsafe.Sizeof(EntryT(0)))
	BITS_PER_ENTRY       = (types.IDX_T)(BYTES_PER_ENTRY) * 8
	MAX_ENTRY            = ^EntryT(0)
	STANDARD_ENTRY_COUNT = (constants.STANDARD_VECTOR_SIZE + (BITS_PER_ENTRY - 1)) / BITS_PER_ENTRY
)

type EntryIndex struct {
	Idx    types.IDX_T
	Offset types.IDX_T
}

type ValidityMask struct {
	Data []EntryT
}
