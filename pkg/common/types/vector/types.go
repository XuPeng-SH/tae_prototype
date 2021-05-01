package vector

import (
	"fmt"
	"tae/pkg/common/types/vbuff"
	"tae/pkg/common/types/vmask"
)

type VectorType uint8

const (
	FLAT_VECTOR VectorType = iota
	CONSTANT_VECTOR
	DICTIONARY_VECTOR
	SEQUENCE_VECTOR
)

func (vt VectorType) String() string {
	switch vt {
	case FLAT_VECTOR:
		return "FLAT"
	case SEQUENCE_VECTOR:
		return "SEQUENCE"
	case DICTIONARY_VECTOR:
		return "DICTIONARY"
	case CONSTANT_VECTOR:
		return "CONSTANT"
	}
	return fmt.Sprintf("Unkown type: %d", vt)
}

type Vector struct {
	Buff vbuff.IVectorBuffer // The main buffer holding the data of the vector
	// ExtraBuff *vbuff.VectorBuffer // The buffer holding extra data of the vector
	Validity *vmask.ValidityMask
	Type     VectorType
}
