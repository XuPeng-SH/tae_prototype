package value

import (
	"tae/pkg/common/types"
)

// type ValueT struct {
// 	Data []byte
// }

type Value struct {
	Type   types.LogicType
	IsNull bool
	// Val    ValueT
	Data []byte
}
