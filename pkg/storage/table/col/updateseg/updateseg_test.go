package updateseg

import (
	"tae/pkg/common/types"
	// "tae/pkg/common/types/constants"
	"github.com/stretchr/testify/assert"
	"tae/pkg/storage/table/col/coldata"
	"testing"
)

func TestBasic(t *testing.T) {
	start1 := types.IDX_0
	count1 := types.IDX_T(1024)
	cdata := coldata.ColumnData{}
	useg := NewUpdateSegment(&cdata, start1, count1)
	t.Log(useg.String())
}

func TestFindSeg(t *testing.T) {
	start1 := types.IDX_0
	count := UPDATE_SEGMENT_MAX_ROWS
	start2 := start1 + count
	useg1 := NewUpdateSegment(nil, start1, count)
	useg2 := NewUpdateSegment(nil, start2, count)
	useg1.Append(useg2)

	fseg1 := useg1.FindSegByVecIdx(WhichVecIdx(start1))
	assert.Equal(t, fseg1, useg1)
	fseg2 := useg2.FindSegByVecIdx(WhichVecIdx(start2))
	assert.Equal(t, fseg2, useg2)
	fseg3 := useg1.FindSegByVecIdx(WhichVecIdx(start2))
	assert.Equal(t, fseg3, useg2)
	fseg4 := useg2.FindSegByVecIdx(WhichVecIdx(start2 + count))
	assert.Equal(t, fseg4, nil)

	fseg5 := useg2.FindSegByVecIdx(WhichVecIdx(start1))
	assert.Equal(t, fseg5, nil)
}
