package types

import ()

type IDX_T uint64
type PhysicalType uint8

const (
	P_NA PhysicalType = iota
	P_BOOL
	P_UINT8
	P_INT8
	P_UINT16
	P_INT16
	P_INVALID
)

type LogicalTypeId uint8

const (
	INVALID LogicalTypeId = iota
	SQLNULL               = 1 /* NULL type, used for constant NULL */
	UNKNOWN               = 2 /* unknown type, used for parameter expressions */
	ANY                   = 3 /* ANY type, used for functions that accept any type as parameter */

	BOOLEAN   = 10
	TINYINT   = 11
	SMALLINT  = 12
	INTEGER   = 13
	BIGINT    = 14
	DATE      = 15
	TIME      = 16
	TIMESTAMP = 17
	DECIMAL   = 18
	FLOAT     = 19
	DOUBLE    = 20
	CHAR      = 21
	VARCHAR   = 22
	BLOB      = 24
	INTERVAL  = 25
	UTINYINT  = 26
	USMALLINT = 27
	UINTEGER  = 28
	UBIGINT   = 29
	HUGEINT   = 50
	POINTER   = 51
	HASH      = 52
	VALIDITY  = 53

	STRUCT = 100
	LIST   = 101
	MAP    = 102
	TABLE  = 103
)

type LogicType struct {
	PType     PhysicalType
	LType     LogicalTypeId
	Width     uint8
	Scale     uint8
	Collation string
}
