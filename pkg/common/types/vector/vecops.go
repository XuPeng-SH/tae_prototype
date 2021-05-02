package vector

import (
	"fmt"
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
	"tae/pkg/common/types/selvec"

	log "github.com/sirupsen/logrus"
)

func CopyWithSelVec(src_vec, des_vec *Vector, src_offset, des_offset, count types.IDX_T, sel_vec *selvec.SelectionVector) {
	if des_vec.Type != FLAT_VECTOR {
		panic(fmt.Sprintf("Should not call CopyWithSelVec for target vector type: %s", src_vec.Type))
	}
	if des_vec.GetBuffer().MaxItems() < des_offset+count {
		panic(fmt.Sprintf("Dest overflow %d/%d", des_vec.GetBuffer().MaxItems(), des_offset+count))
	}
	if src_vec.GetLogicType() != des_vec.GetLogicType() {
		msg := fmt.Sprintf("Src %s and dest %s type mismatch", src_vec.GetLogicType(), des_vec.GetLogicType())
		log.Error(msg)
		panic(msg)
	}
	if count == 0 {
		return
	}
	active_sv := sel_vec
	switch src_vec.Type {
	case CONSTANT_VECTOR:
		active_sv = selvec.ZERO_SV
	case DICTIONARY_VECTOR:
		// PXU TODO
	case FLAT_VECTOR:
	default:
		panic(fmt.Sprintf("Should not call CopyWithSelVec for vector type: %v", src_vec.Type))
	}

	// Handle mask
	tmask := des_vec.GetValidity()
	if src_vec.Type == CONSTANT_VECTOR {
		if src_vec.IsNull() {
			for i := types.IDX_0; i < count; i++ {
				tmask.SetInvalid(i)
			}
		}
	} else {
		smask := src_vec.GetValidity()
		if !smask.AllValid() {
			for i := types.IDX_0; i < count; i++ {
				index := types.IDX_T(active_sv.GetIndex(i))
				valid := smask.IsRowValid(index)
				if valid {
					tmask.SetValid(index)
				} else {
					tmask.SetInvalid(index)
				}
			}
		}
	}

	// Handle data
	data_size := src_vec.GetLogicType().GetPhysicalType().Size()
	src_data := src_vec.GetBuffer().GetData()
	des_data := des_vec.GetBuffer().GetData()
	for i := types.IDX_0; i < count; i++ {
		index := types.IDX_T(active_sv.GetIndex(src_offset + i))
		copy(des_data[(des_offset+i)*data_size:], src_data[(src_offset+index)*data_size:(src_offset+index+1)*data_size])
	}
}

func Copy(src_vec, des_vec *Vector, src_offset, des_offset, count types.IDX_T) {
	switch src_vec.Type {
	case CONSTANT_VECTOR:
		CopyWithSelVec(src_vec, des_vec, src_offset, des_offset, count, selvec.ZERO_SV)
		return
	case DICTIONARY_VECTOR:
		// PXU TODO
		return
	case FLAT_VECTOR:
		sel := selvec.SEQUENTIAL_SV
		if des_offset+count > constants.STANDARD_VECTOR_SIZE {
			sel = selvec.New(selvec.WithCount(des_offset + count))
		}
		CopyWithSelVec(src_vec, des_vec, src_offset, des_offset, count, sel)
		return
	}
	panic(fmt.Sprintf("Should not call Copy for source vector type: %v", src_vec.Type))
}
