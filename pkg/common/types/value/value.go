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
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.SQLNULL)))
		return val
	}
	switch x := vs[0].(type) {
	case nil:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.SQLNULL)))
	case types.LogicType:
		val.Type = x
	case bool:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.BOOLEAN)))
		val.Val.Data = make([]byte, 1)
		if x {
			val.Val.Data[0] = byte(1)
		} else {
			val.Val.Data[0] = byte(0)
		}
		val.IsNull = false
	case int8:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.TINYINT)))
		val.Val.Data = make([]byte, 1)
		val.Val.Data[0] = byte(x)
		val.IsNull = false
	case uint8:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.UTINYINT)))
		val.Val.Data = make([]byte, 1)
		val.Val.Data[0] = byte(x)
		val.IsNull = false
	case int16:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.SMALLINT)))
		val.Val.Data = make([]byte, 2)
		binary.BigEndian.PutUint16(val.Val.Data, uint16(x))
		val.IsNull = false
	case uint16:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.USMALLINT)))
		val.Val.Data = make([]byte, 2)
		binary.BigEndian.PutUint16(val.Val.Data, uint16(x))
		val.IsNull = false
	case int32:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.INTEGER)))
		val.Val.Data = make([]byte, 4)
		binary.BigEndian.PutUint32(val.Val.Data, uint32(x))
		val.IsNull = false
	case uint32:
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.UINTEGER)))
		val.Val.Data = make([]byte, 4)
		binary.BigEndian.PutUint32(val.Val.Data, uint32(x))
		val.IsNull = false
	case int64:
		val.Val.Data = make([]byte, 8)
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.BIGINT)))
		binary.BigEndian.PutUint64(val.Val.Data, uint64(x))
		val.IsNull = false
	case uint64:
		val.Val.Data = make([]byte, 8)
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.UBIGINT)))
		binary.BigEndian.PutUint64(val.Val.Data, uint64(x))
		val.IsNull = false
	case int:
		val.Val.Data = make([]byte, 8)
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.BIGINT)))
		binary.BigEndian.PutUint64(val.Val.Data, uint64(x))
		val.IsNull = false
	case float32:
		val.Val.Data = make([]byte, 4)
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.FLOAT32)))
		binary.BigEndian.PutUint32(val.Val.Data, math.Float32bits(x))
		val.IsNull = false
	case float64:
		val.Val.Data = make([]byte, 8)
		val.Type = *(types.NewLogicType(types.WithLogicTypeId(types.FLOAT64)))
		binary.BigEndian.PutUint64(val.Val.Data, math.Float64bits(x))
		val.IsNull = false
	default:
		panic(fmt.Sprintf("UNKNOWN value type %T", x))
	}

	return val
}

func (v *Value) GetValue() (ret interface{}) {
	if v.IsNull {
		return ret
	}
	switch v.Type.LType {
	case types.BOOLEAN:
		if v.Val.Data[0] == 0 {
			ret = false
		} else {
			ret = true
		}
	case types.TINYINT:
		ret = int8(v.Val.Data[0])
	case types.UTINYINT:
		ret = uint8(v.Val.Data[0])
	case types.SMALLINT:
		ret = (int16)(binary.BigEndian.Uint16(v.Val.Data))
	case types.USMALLINT:
		ret = binary.BigEndian.Uint16(v.Val.Data)
	case types.INTEGER:
		ret = (int32)(binary.BigEndian.Uint32(v.Val.Data))
	case types.UINTEGER:
		ret = binary.BigEndian.Uint32(v.Val.Data)
	case types.BIGINT:
		ret = (int64)(binary.BigEndian.Uint64(v.Val.Data))
	case types.UBIGINT:
		ret = binary.BigEndian.Uint64(v.Val.Data)
	case types.FLOAT32:
		bytes := binary.BigEndian.Uint32(v.Val.Data)
		ret = math.Float32frombits(bytes)
	case types.FLOAT64:
		bytes := binary.BigEndian.Uint64(v.Val.Data)
		ret = math.Float64frombits(bytes)
	default:
		panic(fmt.Sprintf("UNKNOWN value type %v", v.Type.LType))
	}
	return ret
}

// func Min(lt LogicType) Value {
// }
