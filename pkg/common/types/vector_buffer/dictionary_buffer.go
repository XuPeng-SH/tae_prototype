package vector_buffer

import (
	"tae/pkg/common/types"
	SV "tae/pkg/common/types/selection_vector"
)

type DictionaryBuffer struct {
	VectorBuffer
	SelVec *SV.SelectionVector
}

func NewDictonaryBuffer(options ...DictBuffOption) *DictionaryBuffer {
	sv := &DictionaryBuffer{
		VectorBuffer: *NewVectorBuffer(WithType(DICTIONARY_BUFFER)),
	}
	for _, option := range options {
		*sv = option(*sv)
	}
	return sv
}

type DictBuffOption func(DictionaryBuffer) DictionaryBuffer

func WithDictBuffCount(count int) DictBuffOption {
	return func(db DictionaryBuffer) DictionaryBuffer {
		db.SelVec = SV.New(SV.WithCount(count))
		return db
	}
}

func WithDictBuffSelectionVector(sv SV.SelectionVector) DictBuffOption {
	return func(db DictionaryBuffer) DictionaryBuffer {
		db.SelVec = SV.New()
		db.SelVec.InitWithOther(sv)
		return db
	}
}

func WithBuffWithItemType(lt types.LogicType) DictBuffOption {
	return func(db DictionaryBuffer) DictionaryBuffer {
		db.VectorBuffer.ItemType = lt
		return db
	}
}
