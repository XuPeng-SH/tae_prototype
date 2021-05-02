package selvec

import (
	"tae/pkg/common/types"
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

type ISelectionVector interface {
	InitWithCount(count types.IDX_T)
	InitWithData(data *SelectionData)
	InitWithOther(other ISelectionVector)
	Empty() bool
	SetIndex(index types.IDX_T, loc EntryT)
	GetIndex(index types.IDX_T) EntryT
	Count() types.IDX_T
	Slice(other SelectionVector, count types.IDX_T) *SelectionData
	Swap(i, j types.IDX_T)
	String() string
	ToString(count types.IDX_T) string
	GetData() *SelectionData
}

type SelectionData struct {
	Data []EntryT
}

type SelectionVector struct {
	Data *SelectionData
}
