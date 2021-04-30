package types

import ()

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

func (lt *LogicType) getPhysicalType() PhysicalType {
	switch lt.LType {
	case BOOLEAN:
		return P_BOOL
	case TINYINT:
		return P_INT8
	}
	return P_INVALID
}
