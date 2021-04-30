package types

import (
	"fmt"
)

func (pt PhysicalType) Size() uint8 {
	switch pt {
	case P_INVALID:
		return PSIZE_INVALID
	case P_BOOL:
		return PSIZE_BOOL
	case P_UINT8:
		return PSIZE_UINT8
	case P_INT8:
		return PSIZE_INT8
	case P_UINT16:
		return PSIZE_UINT16
	case P_INT16:
		return PSIZE_INT16
	case P_UINT32:
		return PSIZE_UINT32
	case P_INT32:
		return PSIZE_INT32
	case P_UINT64:
		return PSIZE_UINT64
	case P_INT64:
		return PSIZE_INT64
	case P_INT:
		return PSIZE_INT
	case P_FLOAT32:
		return PSIZE_FLOAT32
	case P_FLOAT64:
		return PSIZE_FLOAT64
	}
	panic(fmt.Sprintf("UNKNOWN physical type: %v", pt))
}

func (pt PhysicalType) ToString() string {
	switch pt {
	case P_INVALID:
		return "INVALID"
	case P_BOOL:
		return "BOOL"
	case P_UINT8:
		return "UINT8"
	case P_INT8:
		return "INT8"
	case P_UINT16:
		return "UINT16"
	case P_INT16:
		return "INT16"
	case P_UINT32:
		return "UINT32"
	case P_INT32:
		return "INT32"
	case P_UINT64:
		return "UINT64"
	case P_INT64:
		return "INT64"
	case P_INT: // TODO: not support int
		return "INT64"
	case P_FLOAT32:
		return "FLOAT32"
	case P_FLOAT64:
		return "FLOAT64"
	}
	panic(fmt.Sprintf("UNKNOWN physical type: %v", pt))
}

func (pt PhysicalType) IsConstantSize() bool {
	if pt >= P_BOOL && pt <= P_FLOAT64 {
		return true
	}
	return false
}

func (pt PhysicalType) IsNumeric() bool {
	if pt >= P_UINT8 && pt <= P_FLOAT64 {
		return true
	}
	return false
}

func (pt PhysicalType) IsIntegral() bool {
	if pt >= P_UINT8 && pt <= P_INT {
		return true
	}
	return false
}

func (pt PhysicalType) IsInteger() bool {
	if pt >= P_UINT8 && pt <= P_INT {
		return true
	}
	return false
}
