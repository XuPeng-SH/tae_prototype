package vector

import (
	"fmt"
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	"tae/pkg/common/types/selvec"
	"tae/pkg/common/types/value"
	"tae/pkg/common/types/vbuff"
	"tae/pkg/common/types/vmask"
)

var (
	_ IVector = (*Vector)(nil)
)

func NewVector(options ...Option) *Vector {
	v := &Vector{
		Buff:     vbuff.NewVectorBuffer(vbuff.WithItemType(types.LT_INVALID)),
		Type:     FLAT_VECTOR,
		Validity: vmask.New(types.IDX_0),
	}
	for _, option := range options {
		*v = option(*v)
	}
	return v
}

type Option func(Vector) Vector

func WithInitByLogicType(lt types.LogicType) Option {
	return func(vec Vector) Vector {
		vec.Type = FLAT_VECTOR
		vec.Buff = vbuff.NewVectorBuffer(vbuff.WithItemType(lt),
			vbuff.WithSize((types.IDX_T)(lt.GetPhysicalType().Size())*constants.STANDARD_VECTOR_SIZE))
		return vec
	}
}

func WithInitByValue(val *value.Value) Option {
	return func(vec Vector) Vector {
		vec.Type = CONSTANT_VECTOR
		lt := val.GetLogicType()
		vec.Buff = vbuff.NewVectorBuffer(vbuff.WithItemType(lt),
			vbuff.WithSize((types.IDX_T)(lt.GetPhysicalType().Size())))
		vec.SetValue(0, val)
		return vec
	}
}

func (vec *Vector) Flatten(count types.IDX_T) {
	switch vec.Type {
	case FLAT_VECTOR:
		return
	case DICTIONARY_VECTOR:
		// TODO
		return
	case CONSTANT_VECTOR:
		if vec.IsNull() {
			vec.Validity.MakeRoom(constants.STANDARD_VECTOR_SIZE)
			vec.Validity.InvalidateRows(count)
			return
		}
		vec.Buff.Resize(false)
		vec.Buff.ForceRepeat(0, count)
		vec.Type = FLAT_VECTOR
		return
	}
	panic(fmt.Sprintf("Should not call Flatten for vector type: %v", vec.Type))
}

func (vec *Vector) ReferenceOther(other IVector, offset types.IDX_T) {
	vec.Type = other.GetType()
	if offset == 0 {
		vec.Buff = other.GetBuffer()
		vec.Validity = other.GetValidity()
		return
	}

	vec.Buff.ReferenceOther(other.GetBuffer(), offset)
	vec.Validity.Slice(*other.GetValidity(), offset)
}

func (vec *Vector) SliceOther(other IVector, offset types.IDX_T) {
	if other.GetType() == CONSTANT_VECTOR {
		vec.ReferenceOther(other, 0)
	}
	if other.GetType() != FLAT_VECTOR || vec.Type != FLAT_VECTOR {
		panic("Slice should only on FLAT_VECTOR")
	}
	vec.ReferenceOther(other, offset)
}

func (vec *Vector) SliceOtherWithSel(other IVector, sel selvec.ISelectionVector, count types.IDX_T) {
	vec.ReferenceOther(other, 0)
	vec.SliceWithSel(sel, count)
}

func (vec *Vector) SliceWithSel(sel selvec.ISelectionVector, count types.IDX_T) {
	if vec.Type == CONSTANT_VECTOR {
		return
	}
	vec.Buff = vbuff.NewDictonaryBuffer(vbuff.WithDictBuffItemType(vec.Buff.GetItemType()),
		vbuff.WithDictBuffSelectionVector(sel))
	vec.Type = DICTIONARY_VECTOR
}

func (vec *Vector) GetBuffer() vbuff.IVectorBuffer {
	return vec.Buff
}

func (vec *Vector) GetType() VectorType {
	return vec.Type
}

func (vec *Vector) GetLogicType() types.LogicType {
	return vec.Buff.GetItemType()
}

func (vec *Vector) SetValue(idx types.IDX_T, val *value.Value) {
	lt := vec.GetLogicType()
	if lt != val.GetLogicType() {
		// PXU TODO: Try Cast
		return
	}

	vec.Buff.SetValue(idx, val.GetData())
}

func (vec *Vector) GetValidity() *vmask.ValidityMask {
	return vec.Validity
}

func (vec *Vector) GetValue(idx types.IDX_T) interface{} {
	switch vec.Type {
	case CONSTANT_VECTOR:
		idx = 0
	}
	if !vec.Validity.IsRowValid(idx) {
		// TODO: Init all these value as const value
		return *(value.NewValue(vec.GetLogicType()))
	}
	return vec.Buff.GetValue(idx)
}

func (vec *Vector) String() string {
	ret := "Vec(" + vec.Type.String() + ")," + vec.Buff.String()
	return ret
}

func (vec *Vector) IsNull(opt ...interface{}) bool {
	switch vec.Type {
	case CONSTANT_VECTOR:
		return !vec.Validity.IsRowValid(0)
	case FLAT_VECTOR:
		idx := opt[0].(types.IDX_T)
		return !vec.Validity.IsRowValid(idx)
	}
	panic(fmt.Sprintf("Should not call IsNull for vector type: %v", vec.Type))
}

func (vec *Vector) SetNull(is_null bool, opt ...interface{}) {
	switch vec.Type {
	case CONSTANT_VECTOR:
		if is_null {
			vec.Validity.SetInvalid(0)
		} else {
			vec.Validity.SetValid(0)
		}
	case FLAT_VECTOR:
		idx := opt[0].(types.IDX_T)
		if is_null {
			vec.Validity.SetInvalid(idx)
		} else {
			vec.Validity.SetValid(idx)
		}
	}
	panic(fmt.Sprintf("Should not call SetNull for vector type: %v", vec.Type))
}
