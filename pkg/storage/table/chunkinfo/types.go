package chunkinfo

import (
	"fmt"
	types "tae/pkg/common/types"
)

type ChunkInfoType uint8

const (
	CONSTANT_INFO ChunkInfoType = iota
	VECTOR_INFO
	EMPTY_INFO
)

func (cit ChunkInfoType) String() string {
	switch cit {
	case CONSTANT_INFO:
		return "CONSTANT_INFO"
	case VECTOR_INFO:
		return "VECTOR_INFO"
	case EMPTY_INFO:
		return "EMPTY_INFO"
	}
	return fmt.Sprintf("Unkown type: %d", cit)
}

type ChunkInfo struct {
	Start types.IDX_T
	Type  ChunkInfoType
}
