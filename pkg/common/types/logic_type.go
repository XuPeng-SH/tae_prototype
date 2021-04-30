package types

import (
	"fmt"
)

var (
	LT_INVALID   = LogicType{LType: INVALID}
	LT_SQLNULL   = LogicType{LType: SQLNULL}
	LT_BOOLEAN   = LogicType{LType: BOOLEAN}
	LT_TINYINT   = LogicType{LType: TINYINT}
	LT_UTINYINT  = LogicType{LType: UTINYINT}
	LT_SMALLINT  = LogicType{LType: SMALLINT}
	LT_USMALLINT = LogicType{LType: USMALLINT}
	LT_INTEGER   = LogicType{LType: INTEGER}
	LT_UINTEGER  = LogicType{LType: UINTEGER}
	LT_BIGINT    = LogicType{LType: BIGINT}
	LT_UBIGINT   = LogicType{LType: UBIGINT}
	LT_FLOAT32   = LogicType{LType: FLOAT32}
	LT_FLOAT64   = LogicType{LType: FLOAT64}
)

func NewLogicType(options ...LTOption) *LogicType {
	lt := &LogicType{
		LType: INVALID,
	}
	for _, option := range options {
		*lt = option(*lt)
	}
	return lt
}

type LTOption func(LogicType) LogicType

func WithLogicTypeId(id LogicalTypeId) LTOption {
	return func(lt LogicType) LogicType {
		lt.LType = id
		return lt
	}
}

func WithWidthAndScale(width, scale uint8) LTOption {
	return func(lt LogicType) LogicType {
		lt.Width = width
		lt.Scale = scale
		return lt
	}
}

func WithCollation(collation string) LTOption {
	return func(lt LogicType) LogicType {
		lt.Collation = collation
		return lt
	}
}

func (lt *LogicType) GetPhysicalType() PhysicalType {
	switch lt.LType {
	case BOOLEAN:
		return P_BOOL
	case TINYINT:
		return P_INT8
	case UTINYINT:
		return P_UINT8
	case SMALLINT:
		return P_INT16
	case USMALLINT:
		return P_UINT16
	case INTEGER:
		return P_INT32
	case UINTEGER:
		return P_UINT32
	case BIGINT:
		return P_INT64
	case UBIGINT:
		return P_UINT64
	case FLOAT32:
		return P_FLOAT32
	case FLOAT64:
		return P_FLOAT64
	case INVALID:
		return P_INVALID
	case UNKNOWN:
		return P_INVALID
	}
	panic(fmt.Sprintf("UNKNOWN type: %v", lt.LType))
}

func (lt *LogicType) GetID() LogicalTypeId {
	return lt.LType
}
