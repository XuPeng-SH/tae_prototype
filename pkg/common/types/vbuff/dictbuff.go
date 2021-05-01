package vbuff

import (
	"tae/pkg/common/types"
	"tae/pkg/common/types/selvec"
)

type DictionaryBuffer struct {
	VectorBuffer
	SelVec *selvec.SelectionVector
}

var (
	_ IVectorBuffer = (*DictionaryBuffer)(nil)
)

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
		db.SelVec = selvec.New(selvec.WithCount(count))
		return db
	}
}

func WithDictBuffSelectionVector(sv selvec.SelectionVector) DictBuffOption {
	return func(db DictionaryBuffer) DictionaryBuffer {
		db.SelVec = selvec.New()
		db.SelVec.InitWithOther(sv)
		return db
	}
}

func WithDictBuffItemType(lt types.LogicType) DictBuffOption {
	return func(db DictionaryBuffer) DictionaryBuffer {
		db.VectorBuffer.ItemType = lt
		return db
	}
}
