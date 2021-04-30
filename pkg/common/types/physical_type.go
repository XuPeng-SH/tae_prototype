package types

import (
	"fmt"
)

func (pt *PhysicalType) Size() uint8 {
	switch *pt {
	case P_NA:
		return PSIZE_NA
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
	panic(fmt.Sprintf("UNKNOWN physical type: %v", *pt))
}
