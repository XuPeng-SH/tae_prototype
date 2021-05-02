package selvec

import (
	// "tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	"unsafe"
)

type EntryT uint32

const (
	EntryBytes = (uint)(unsafe.Sizeof(EntryT(0)))
)

var (
	ZERO_SV = New(WithCount(constants.STANDARD_VECTOR_SIZE))
)

type SelectionData struct {
	Data []EntryT
}

type SelectionVector struct {
	Data *SelectionData
}
