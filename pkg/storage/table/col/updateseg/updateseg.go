package updateseg

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	"tae/pkg/storage/table/col/coldata"
	seg "tae/pkg/storage/table/segment"
)

var (
	_ seg.ISegment = (*UpdateSegment)(nil)
)

func WhichVecIdx(row_id types.IDX_T) types.IDX_T {
	return row_id / constants.STANDARD_VECTOR_SIZE
}

func NewUpdateSegment(data *coldata.ColumnData, start, count types.IDX_T) *UpdateSegment {
	useg := &UpdateSegment{
		Data:     data,
		ISegment: seg.NewSegment(start, count),
	}
	return useg
}

func (useg *UpdateSegment) String() string {
	return fmt.Sprintf("USeg(%s,%s)", useg.Data, useg.ISegment.String())
}

func (useg *UpdateSegment) HasUpdate() bool {
	if useg.Tree == nil {
		return false
	}
	return true
}

func (useg *UpdateSegment) Capacity() types.IDX_T {
	return UPDATE_SEGMENT_MAX_ROWS
}

func (useg *UpdateSegment) HasVectorUpdate(vec_idx types.IDX_T) bool {
	if !useg.HasUpdate() {
		return false
	}
	// PXU TODO: Check UpdateInfo
	return true
}

func (useg *UpdateSegment) HasVectorRangeUpdate(vstart_idx, vend_idx types.IDX_T) bool {
	global_vec_index := WhichVecIdx(useg.GetStartRow())
	if vstart_idx < global_vec_index {
		panic(fmt.Sprintf("Start vector index should not be less than: %d", global_vec_index))
	}
	active_useg := useg
	for i := vstart_idx; i < vend_idx; i++ {
		index := i - global_vec_index
		for index >= MOSEL_VECTOR_COUNT {
			active_useg = useg.GetNext().(*UpdateSegment)
		}
		if active_useg.HasVectorUpdate(index) {
			return true
		}
	}
	return false
}

func (useg *UpdateSegment) FindSegByVecIdx(vec_idx types.IDX_T) seg.ISegment {
	base_vec_index := WhichVecIdx(useg.GetStartRow())
	if vec_idx < base_vec_index {
		log.Warnf(fmt.Sprintf("Start vector index %d should not be less than: %d", vec_idx, base_vec_index))
		// panic(fmt.Sprintf("Start vector index should not be less than: %d", base_vec_index))
		return nil
	}
	var active_useg seg.ISegment
	active_useg = useg
	for vec_idx >= MOSEL_VECTOR_COUNT+base_vec_index {
		active_useg = useg.GetNext()
		if active_useg == nil {
			return nil
		}
		base_vec_index += MOSEL_VECTOR_COUNT
	}
	return active_useg
}
