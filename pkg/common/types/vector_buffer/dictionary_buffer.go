package vector_buffer

import (
	"tae/pkg/common/types/selection_vector"
)

type DictionaryBuffer struct {
	VectorBuffer
	SelVec selection_vector.SelectionVector
}
