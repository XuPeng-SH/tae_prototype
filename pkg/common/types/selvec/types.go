package selvec

import (
	"unsafe"
)

type EntryT uint32

const (
	EntryBytes = (uint)(unsafe.Sizeof(EntryT(0)))
)

type SelectionData struct {
	Data []EntryT
}

type SelectionVector struct {
	Data *SelectionData
}
