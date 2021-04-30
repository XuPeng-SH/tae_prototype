package vector_buffer

import (
	"encoding/binary"
	"fmt"
	"math"
	"tae/pkg/common/types"
)

func NewVectorBuffer(options ...Option) *VectorBuffer {
	vf := &VectorBuffer{
		Type: STANDARD_BUFFER,
	}
	for _, option := range options {
		*vf = option(*vf)
	}
	vf.ItemSize = vf.ItemType.GetPhysicalType().Size()
	return vf
}

type Option func(VectorBuffer) VectorBuffer

func WithSize(size int) Option {
	return func(vf VectorBuffer) VectorBuffer {
		if size < 0 {
			panic("")
		}
		vf.Data = make([]byte, size)
		return vf
	}
}

func WithType(t VectorBufferType) Option {
	return func(vf VectorBuffer) VectorBuffer {
		vf.Type = t
		return vf
	}
}

func WithItemType(lt types.LogicType) Option {
	return func(vf VectorBuffer) VectorBuffer {
		vf.ItemType = lt
		return vf
	}
}

func (vf *VectorBuffer) GetType() VectorBufferType {
	return vf.Type
}

func (vf *VectorBuffer) GetItemType() types.LogicType {
	return vf.ItemType
}

func (vf *VectorBuffer) Size() int {
	return len(vf.Data)
}

func (vf *VectorBuffer) GetValue(idx int, val interface{}) {
}

func (vf *VectorBuffer) SetValue(idx int, val interface{}) {
	if idx < 0 || (idx+1)*(int)(vf.ItemSize) > len(vf.Data) {
		panic(fmt.Sprintf("Invalid idx: %d, should be in range [0, %d)", idx, len(vf.Data)/(int)(vf.ItemSize)))
	}
	switch vf.ItemType.GetID() {
	case types.BOOLEAN:
		vf.Data[idx] = val.(byte)
		return
	case types.TINYINT:
		vf.Data[idx] = val.(byte)
		return
	case types.UTINYINT:
		vf.Data[idx] = val.(byte)
		return
	case types.SMALLINT:
		binary.BigEndian.PutUint16(vf.Data[(idx*(int)(vf.ItemSize)):], (uint16)(val.(int16)))
		return
	case types.USMALLINT:
		binary.BigEndian.PutUint16(vf.Data[(idx*(int)(vf.ItemSize)):], val.(uint16))
		return
	case types.INTEGER:
		binary.BigEndian.PutUint32(vf.Data[(idx*(int)(vf.ItemSize)):], (uint32)(val.(int32)))
		return
	case types.UINTEGER:
		binary.BigEndian.PutUint32(vf.Data[(idx*(int)(vf.ItemSize)):], val.(uint32))
		return
	case types.BIGINT:
		binary.BigEndian.PutUint64(vf.Data[(idx*(int)(vf.ItemSize)):], (uint64)(val.(int64)))
		return
	case types.UBIGINT:
		binary.BigEndian.PutUint64(vf.Data[(idx*(int)(vf.ItemSize)):], val.(uint64))
		return
	case types.FLOAT32:
		f := val.(float32)
		binary.BigEndian.PutUint32(vf.Data[(idx*(int)(vf.ItemSize)):], math.Float32bits(f))
		return
	case types.FLOAT64:
		f := val.(float64)
		binary.BigEndian.PutUint64(vf.Data[(idx*(int)(vf.ItemSize)):], math.Float64bits(f))
		return
	}
	panic(fmt.Sprintf("UNIMPLEMENTED logic type: %v", vf.ItemType))
}
