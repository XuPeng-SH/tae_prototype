package selection_vector

import (
	"bytes"
	"unsafe"
)

type EntryT uint32

const (
	EntryBytes = (uint)(unsafe.Sizeof(EntryT(0)))
)

type SelectionVector struct {
	Data *bytes.Buffer
}
