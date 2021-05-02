package vbuff

import (
	"encoding/binary"
	"fmt"
	"math"
	"tae/pkg/common/types"

	log "github.com/sirupsen/logrus"
)

var (
	_ IVectorBuffer = (*VectorBuffer)(nil)
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

func WithSize(size types.IDX_T) Option {
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

func (vf *VectorBuffer) Size() types.IDX_T {
	return (types.IDX_T)(len(vf.Data))
}

func (vf *VectorBuffer) GetValue(idx types.IDX_T) (ret interface{}) {
	if (idx+1)*(types.IDX_T)(vf.ItemSize) > (types.IDX_T)(len(vf.Data)) {
		msg := fmt.Sprintf("Invalid idx: %d, should be in range [0, %d)", idx, len(vf.Data)/(int)(vf.ItemSize))
		log.Error(msg)
		panic(msg)
	}

	switch vf.ItemType.GetID() {
	case types.BOOLEAN:
		if vf.Data[idx] > 0 {
			ret = true
		} else {
			ret = false
		}
	case types.TINYINT:
		ret = (int8)(vf.Data[idx])
	case types.UTINYINT:
		ret = (uint8)(vf.Data[idx])
	case types.SMALLINT:
		ret = (int16)(binary.BigEndian.Uint16(vf.Data[(idx * vf.ItemSize):]))
	case types.USMALLINT:
		ret = binary.BigEndian.Uint16(vf.Data[(idx * vf.ItemSize):])
	case types.INTEGER:
		ret = (int32)(binary.BigEndian.Uint32(vf.Data[(idx * vf.ItemSize):]))
	case types.UINTEGER:
		ret = binary.BigEndian.Uint32(vf.Data[(idx * vf.ItemSize):])
	case types.BIGINT:
		ret = (int64)(binary.BigEndian.Uint64(vf.Data[(idx * vf.ItemSize):]))
	case types.UBIGINT:
		ret = binary.BigEndian.Uint64(vf.Data[(idx * vf.ItemSize):])
	case types.FLOAT32:
		bytes := binary.BigEndian.Uint32(vf.Data[(idx * vf.ItemSize):])
		ret = math.Float32frombits(bytes)
	case types.FLOAT64:
		bytes := binary.BigEndian.Uint64(vf.Data[(idx * vf.ItemSize):])
		ret = math.Float64frombits(bytes)
	default:
		panic(fmt.Sprintf("UNIMPLEMENTED logic type: %v", vf.ItemType))
	}
	return ret
}

func (vf *VectorBuffer) SetValue(idx types.IDX_T, val interface{}) {
	if (idx+1)*vf.ItemSize > (types.IDX_T)(len(vf.Data)) {
		panic(fmt.Sprintf("Invalid idx: %d, should be in range [0, %d)", idx, len(vf.Data)/(int)(vf.ItemSize)))
	}
	switch bytes := val.(type) {
	case []byte:
		copy(vf.Data[idx:], bytes)
		return
	}
	switch vf.ItemType.GetID() {
	case types.BOOLEAN:
		v := val.(bool)
		if v {
			vf.Data[idx] = byte(1)
		} else {
			vf.Data[idx] = byte(0)
		}
	case types.TINYINT:
		vf.Data[idx] = byte(val.(int8))
	case types.UTINYINT:
		vf.Data[idx] = byte(val.(uint8))
	case types.SMALLINT:
		binary.BigEndian.PutUint16(vf.Data[(idx*vf.ItemSize):], (uint16)(val.(int16)))
	case types.USMALLINT:
		binary.BigEndian.PutUint16(vf.Data[(idx*vf.ItemSize):], val.(uint16))
	case types.INTEGER:
		binary.BigEndian.PutUint32(vf.Data[(idx*vf.ItemSize):], (uint32)(val.(int32)))
	case types.UINTEGER:
		binary.BigEndian.PutUint32(vf.Data[(idx*vf.ItemSize):], val.(uint32))
	case types.BIGINT:
		binary.BigEndian.PutUint64(vf.Data[(idx*vf.ItemSize):], (uint64)(val.(int64)))
	case types.UBIGINT:
		binary.BigEndian.PutUint64(vf.Data[(idx*vf.ItemSize):], val.(uint64))
	case types.FLOAT32:
		f := val.(float32)
		binary.BigEndian.PutUint32(vf.Data[(idx*vf.ItemSize):], math.Float32bits(f))
	case types.FLOAT64:
		f := val.(float64)
		binary.BigEndian.PutUint64(vf.Data[(idx*vf.ItemSize):], math.Float64bits(f))
	default:
		panic(fmt.Sprintf("UNIMPLEMENTED logic type: %v", vf.ItemType))
	}
}

func (vb *VectorBuffer) GetData() []byte {
	return vb.Data
}

func (vb *VectorBuffer) MaxItems() types.IDX_T {
	return (types.IDX_T)(len(vb.Data)) / vb.ItemSize
}

func (vb *VectorBuffer) GetItemSize() types.IDX_T {
	return vb.ItemSize
}

func (vb *VectorBuffer) String() string {
	count := vb.MaxItems()
	if count >= 32 {
		count = 32
	}
	return vb.ToString(count)
}

func (vb *VectorBuffer) ToString(opts ...interface{}) string {
	count := vb.MaxItems()
	if len(opts) > 0 {
		cnt := opts[0].(types.IDX_T)
		if cnt <= 0 {
			count = 0
		} else if cnt < count {
			count = cnt
		}
	}
	ret := "Vbuff(" + vb.Type.String() + "," + vb.ItemType.String() + "): [" + count.String()
	ret += "/" + vb.MaxItems().String() + "] = ["
	for i := types.IDX_0; i < count; i++ {
		ret += fmt.Sprintf("%v", vb.GetValue(i))
		if i != count-1 {
			ret += ","
		}
	}
	ret += "]"
	return ret
}

func (vb *VectorBuffer) ReferenceOther(other IVectorBuffer, offset types.IDX_T) {
	if offset >= vb.MaxItems() {
		panic(fmt.Sprintf("offset %d should be in [%d, %d)", offset, 0, vb.MaxItems()))
	}
	vb.Type = other.GetType()
	vb.ItemType = other.GetItemType()
	vb.ItemSize = other.GetItemSize()
	vb.Data = other.GetData()[offset*vb.ItemSize:]
}
