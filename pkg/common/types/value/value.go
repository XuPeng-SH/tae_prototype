package value

import (
	"encoding/binary"
	// log "github.com/sirupsen/logrus"
	"fmt"
	"math"
	"tae/pkg/common/types"
)

func NewValue(vs ...interface{}) *Value {
	val := &Value{IsNull: true}
	if len(vs) == 0 {
		val.Type = types.LT_SQLNULL
		return val
	}
	switch x := vs[0].(type) {
	case nil:
		val.Type = types.LT_SQLNULL
	case types.LogicType:
		val.Type = x
	case bool:
		val.Type = types.LT_BOOLEAN
		val.Data = make([]byte, 1)
		if x {
			val.Data[0] = byte(1)
		} else {
			val.Data[0] = byte(0)
		}
		val.IsNull = false
	case int8:
		val.Type = types.LT_TINYINT
		val.Data = make([]byte, 1)
		val.Data[0] = byte(x)
		val.IsNull = false
	case uint8:
		val.Type = types.LT_UTINYINT
		val.Data = make([]byte, 1)
		val.Data[0] = byte(x)
		val.IsNull = false
	case int16:
		val.Type = types.LT_SMALLINT
		val.Data = make([]byte, 2)
		binary.BigEndian.PutUint16(val.Data, uint16(x))
		val.IsNull = false
	case uint16:
		val.Type = types.LT_USMALLINT
		val.Data = make([]byte, 2)
		binary.BigEndian.PutUint16(val.Data, uint16(x))
		val.IsNull = false
	case int32:
		val.Type = types.LT_INTEGER
		val.Data = make([]byte, 4)
		binary.BigEndian.PutUint32(val.Data, uint32(x))
		val.IsNull = false
	case uint32:
		val.Type = types.LT_UINTEGER
		val.Data = make([]byte, 4)
		binary.BigEndian.PutUint32(val.Data, uint32(x))
		val.IsNull = false
	case int64:
		val.Data = make([]byte, 8)
		val.Type = types.LT_BIGINT
		binary.BigEndian.PutUint64(val.Data, uint64(x))
		val.IsNull = false
	case uint64:
		val.Data = make([]byte, 8)
		val.Type = types.LT_UBIGINT
		binary.BigEndian.PutUint64(val.Data, uint64(x))
		val.IsNull = false
	case int: // TODO: Should not support int
		val.Data = make([]byte, 8)
		val.Type = types.LT_BIGINT
		binary.BigEndian.PutUint64(val.Data, uint64(x))
		val.IsNull = false
	case float32:
		val.Data = make([]byte, 4)
		val.Type = types.LT_FLOAT32
		binary.BigEndian.PutUint32(val.Data, math.Float32bits(x))
		val.IsNull = false
	case float64:
		val.Data = make([]byte, 8)
		val.Type = types.LT_FLOAT64
		binary.BigEndian.PutUint64(val.Data, math.Float64bits(x))
		val.IsNull = false
	default:
		panic(fmt.Sprintf("UNKNOWN value type %T", x))
	}

	return val
}

func (v *Value) GetData() []byte {
	return v.Data
}

func (v *Value) GetValue() (ret interface{}) {
	if v.IsNull {
		return ret
	}
	switch v.Type.LType {
	case types.BOOLEAN:
		if v.Data[0] == 0 {
			ret = false
		} else {
			ret = true
		}
	case types.TINYINT:
		ret = int8(v.Data[0])
	case types.UTINYINT:
		ret = uint8(v.Data[0])
	case types.SMALLINT:
		ret = (int16)(binary.BigEndian.Uint16(v.Data))
	case types.USMALLINT:
		ret = binary.BigEndian.Uint16(v.Data)
	case types.INTEGER:
		ret = (int32)(binary.BigEndian.Uint32(v.Data))
	case types.UINTEGER:
		ret = binary.BigEndian.Uint32(v.Data)
	case types.BIGINT:
		ret = (int64)(binary.BigEndian.Uint64(v.Data))
	case types.UBIGINT:
		ret = binary.BigEndian.Uint64(v.Data)
	case types.FLOAT32:
		bytes := binary.BigEndian.Uint32(v.Data)
		ret = math.Float32frombits(bytes)
	case types.FLOAT64:
		bytes := binary.BigEndian.Uint64(v.Data)
		ret = math.Float64frombits(bytes)
	default:
		panic(fmt.Sprintf("UNKNOWN value type %v", v.Type.LType))
	}
	return ret
}

func (val *Value) GetLogicType() types.LogicType {
	return val.Type
}

func (val *Value) GetPhysicalTypeSize() uint8 {
	return val.Type.GetPhysicalType().Size()
}

func (val *Value) Clone() *Value {
	ret := &Value{
		Type:   val.Type,
		IsNull: val.IsNull,
		Data:   make([]byte, len(val.Data)),
	}
	copy(ret.Data, val.Data)
	return ret
}

func (val *Value) Cast(lt types.LogicType, strict bool) *Value {
	if val.Type == lt {
		return val.Clone()
	}
	return nil
}

func (val *Value) ToString() string {
	if val.IsNull {
		return fmt.Sprintf("Val[%s](NULL)", val.Type.ToString())
	}
	return fmt.Sprintf("Val[%s](%v)", val.Type.ToString(), val.GetValue())
}
