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

type IValidityMask interface {
	InitAllValid()
	MakeRoom(count types.IDX_T)
	Len() types.IDX_T
	Reset()
	GetEntry(entry_idx types.IDX_T) EntryT
	SetInvalid(row_idx types.IDX_T)
	ValidateRows(rows types.IDX_T)
	InvalidateRows(rows types.IDX_T)
	SetValid(row_idx types.IDX_T)
	IsRowValid(row_idx types.IDX_T) bool
	AllValid() bool
	Slice(other IValidityMask, offset types.IDX_T)
	Combine(other IValidityMask, count types.IDX_T)
	String(count types.IDX_T) string
	GetData() []EntryT
}

type ValidityMask struct {
	Data []EntryT
}
