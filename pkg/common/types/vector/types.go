package vector

import (
	"fmt"
	"tae/pkg/common/types"
	"tae/pkg/common/types/selvec"
	"tae/pkg/common/types/value"
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

type SequenceData struct {
	Start int64
	Step  int64
}

func (sd *SequenceData) String() string {
	return fmt.Sprintf("SeqData(%d,%d)", sd.Start, sd.Step)
}

type ISequenceVector interface {
	GetSequence() (result *SequenceData)
}

type IVector interface {
	ReferenceOther(other IVector, offset types.IDX_T)
	Flatten(count types.IDX_T)
	SliceOther(other IVector, offset types.IDX_T)
	SliceOtherWithSel(other IVector, sel selvec.ISelectionVector, count types.IDX_T)
	SliceWithSel(sel selvec.ISelectionVector, count types.IDX_T)
	GetBuffer() vbuff.IVectorBuffer
	GetType() VectorType
	GetLogicType() types.LogicType
	SetValue(idx types.IDX_T, val *value.Value)
	GetValue(idx types.IDX_T) interface{}
	GetValidity() vmask.IValidityMask
	String() string
	IsNull(opt ...interface{}) bool
	SetNull(is_null bool, opt ...interface{})
	Reset()
	ToSeqenceVector(seq *SequenceData)
}

type Vector struct {
	Buff vbuff.IVectorBuffer // The main buffer holding the data of the vector
	// ExtraBuff *vbuff.VectorBuffer // The buffer holding extra data of the vector
	Validity vmask.IValidityMask
	Type     VectorType
}
