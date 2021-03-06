package types

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type IDX_T = uint64
type SMIDX_T = uint16
type PhysicalType uint8

const (
	IDX_0     = IDX_T(0)
	IDX_1     = IDX_T(1)
	IDX_2     = IDX_T(2)
	IDX_MAX   = ^IDX_0
	SMIDX_0   = SMIDX_T(0)
	SMIDX_1   = SMIDX_T(1)
	SMIDX_2   = SMIDX_T(2)
	SMIDX_MAX = ^SMIDX_0

	UINT32_MAX = ^uint32(0)
	INT32_MAX  = UINT32_MAX >> 1
)

func IDXtoa(idx IDX_T) string {
	return fmt.Sprintf("%d", idx)
}

func AtomicLoad(addr *IDX_T) IDX_T {
	return atomic.LoadUint64(addr)
}

func AtomicStore(addr *IDX_T, val IDX_T) {
	atomic.StoreUint64(addr, val)
}

func AtomicAdd(addr *IDX_T, delta int64) IDX_T {
	return atomic.AddUint64(addr, ^uint64(-delta-1))
}

func AtomicCAS(addr *IDX_T, old, new IDX_T) bool {
	return atomic.CompareAndSwapUint64(addr, old, new)
}

const (
	P_NA PhysicalType = iota
	P_BOOL
	P_BIT
	P_UINT8
	P_INT8
	P_UINT16
	P_INT16
	P_UINT32
	P_INT32
	P_UINT64
	P_INT64
	P_INT
	P_FLOAT32
	P_FLOAT64
	P_INVALID
)

const (
	PSIZE_NA      = IDX_T(0)
	PSIZE_BOOL    = (IDX_T)(unsafe.Sizeof(false))
	PSIZE_UINT8   = (IDX_T)(unsafe.Sizeof(uint8(0)))
	PSIZE_INT8    = (IDX_T)(unsafe.Sizeof(int8(0)))
	PSIZE_UINT16  = (IDX_T)(unsafe.Sizeof(uint16(0)))
	PSIZE_INT16   = (IDX_T)(unsafe.Sizeof(int16(0)))
	PSIZE_UINT32  = (IDX_T)(unsafe.Sizeof(uint32(0)))
	PSIZE_INT32   = (IDX_T)(unsafe.Sizeof(int32(0)))
	PSIZE_UINT64  = (IDX_T)(unsafe.Sizeof(uint64(0)))
	PSIZE_INT64   = (IDX_T)(unsafe.Sizeof(int64(0)))
	PSIZE_INT     = (IDX_T)(unsafe.Sizeof(int(0)))
	PSIZE_FLOAT32 = (IDX_T)(unsafe.Sizeof(float32(0)))
	PSIZE_FLOAT64 = (IDX_T)(unsafe.Sizeof(float64(0)))
	PSIZE_BIT     = IDX_T(1)
	PSIZE_INVALID = IDX_T(0)
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
	FLOAT32   = 19
	FLOAT64   = 20
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
