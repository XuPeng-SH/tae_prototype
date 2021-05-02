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
	ZERO_SV       = New(WithCount(constants.STANDARD_VECTOR_SIZE))
	SEQUENTIAL_SV *SelectionVector
)

func init() {
	sdata := &SelectionData{
		Data: make([]EntryT, constants.STANDARD_VECTOR_SIZE),
	}
	for i := EntryT(0); i < EntryT(constants.STANDARD_VECTOR_SIZE); i++ {
		sdata.Data[i] = i
	}
	SEQUENTIAL_SV = New()
	SEQUENTIAL_SV.InitWithData(sdata)
}

type SelectionData struct {
	Data []EntryT
}

type SelectionVector struct {
	Data *SelectionData
}
