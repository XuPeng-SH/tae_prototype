package vector

import (
	"tae/pkg/common/types"
	"tae/pkg/common/types/value"
	// vmask "tae/pkg/common/types/validity_mask"
	// "fmt"
	svec "tae/pkg/common/types/selection_vector"
	vbuff "tae/pkg/common/types/vector_buffer"
)

func NewVector(options ...Option) *Vector {
	v := &Vector{
		Buff: vbuff.NewVectorBuffer(vbuff.WithItemType(types.LT_INVALID)),
		Type: FLAT_VECTOR,
	}
	for _, option := range options {
		*v = option(*v)
	}
	return v
}

type Option func(Vector) Vector

func WithInitByValue(val value.Value) Option {
	return func(vec Vector) Vector {
		vec.Type = CONSTANT_VECTOR
		lt := val.GetLogicType()
		vec.Buff = vbuff.NewVectorBuffer(vbuff.WithItemType(lt),
			vbuff.WithSize((int)(lt.GetPhysicalType().Size())))
		vec.SetValue(0, val)
		return vec
	}
}

func (vec *Vector) ReferenceOther(other Vector, offset int) {
	vec.Type = other.Type
	if offset == 0 {
		vec.Buff = other.Buff
		vec.Validity = other.Validity
		return
	}

	vec.Buff.ReferenceOther(other.Buff, offset)
	vec.Validity.Slice(*other.Validity, offset)
}

func (vec *Vector) SliceOther(other Vector, offset int) {
	if other.Type == CONSTANT_VECTOR {
		vec.ReferenceOther(other, 0)
	}
	if other.Type != FLAT_VECTOR || vec.Type != FLAT_VECTOR {
		panic("Slice should only on FLAT_VECTOR")
	}
	vec.ReferenceOther(other, offset)
}

func (vec *Vector) SliceOtherWithSel(other Vector, sel svec.SelectionVector, count int) {
	vec.ReferenceOther(other, 0)
	vec.SliceWithSel(sel, count)
}

func (vec *Vector) SliceWithSel(sel svec.SelectionVector, count int) {
	if vec.Type == CONSTANT_VECTOR {
		return
	}
	// PXU TODO
}

func (vec *Vector) GetLogicType() types.LogicType {
	return vec.Buff.GetItemType()
}

func (vec *Vector) SetValue(idx int, val value.Value) {
	lt := vec.GetLogicType()
	if lt != val.GetLogicType() {
		// PXU TODO: Try Cast
		return
	}

	vec.Buff.SetValue(idx, val.GetData())
}
