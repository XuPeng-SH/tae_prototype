package selvec

import (
	"fmt"
	"strconv"
	"tae/pkg/common/types"
)

var (
	_ ISelectionVector = (*SelectionVector)(nil)
)

func New(options ...Option) *SelectionVector {
	sv := &SelectionVector{
		// Data: &SelectionData{},
	}
	for _, option := range options {
		*sv = option(*sv)
	}
	return sv
}

type Option func(SelectionVector) SelectionVector

func WithCount(count types.IDX_T) Option {
	return func(sv SelectionVector) SelectionVector {
		sv.InitWithCount(count)
		return sv
	}
}

func (sv *SelectionVector) InitWithCount(count types.IDX_T) {
	data := &SelectionData{}
	data.Data = make([]EntryT, count)
	sv.Data = data
}

func (sv *SelectionVector) InitWithData(data *SelectionData) {
	sv.Data = data
}

func (sv *SelectionVector) GetData() *SelectionData {
	return sv.Data
}

func (sv *SelectionVector) InitWithOther(other ISelectionVector) {
	sv.Data = other.GetData()
}

func (sv *SelectionVector) Empty() bool {
	return (sv.Data == nil) || (len(sv.Data.Data) == 0)
}

func (sv *SelectionVector) SetIndex(index types.IDX_T, loc EntryT) {
	sv.Data.Data[index] = loc
}

func (sv *SelectionVector) GetIndex(index types.IDX_T) EntryT {
	return sv.Data.Data[index]
}

func (sv *SelectionVector) Count() types.IDX_T {
	if sv.Data == nil {
		return 0
	}
	return types.IDX_T(len(sv.Data.Data))
}

func (sv *SelectionVector) Slice(other SelectionVector, count types.IDX_T) *SelectionData {
	if count > other.Count() {
		return nil
	}
	data := SelectionData{
		Data: make([]EntryT, count),
	}
	for i := types.IDX_0; i < count; i++ {
		new_idx := other.GetIndex(i)
		if (types.IDX_T)(new_idx) > sv.Count() {
			return nil
		}
		idx := sv.GetIndex((types.IDX_T)(new_idx))
		data.Data[i] = idx
	}
	return &data
}

func (sv *SelectionVector) Swap(i, j types.IDX_T) {
	if i == j {
		return
	}
	tmp := sv.GetIndex(i)
	sv.SetIndex((types.IDX_T)(i), sv.GetIndex(j))
	sv.SetIndex((types.IDX_T)(j), tmp)
}

func (sv *SelectionVector) String() string {
	count := sv.Count()
	if count >= 10 {
		count = 10
	}
	return sv.ToString(count)
}

func (sv *SelectionVector) ToString(count types.IDX_T) string {
	ret := fmt.Sprintf("SelectionVector [%v/%v]", count, sv.Count())
	if count > sv.Count() {
		count = sv.Count()
	}
	for i := types.IDX_0; i < count; i++ {
		if i != 0 {
			ret += ", "
		}
		ret += strconv.Itoa((int)(sv.GetIndex(i)))
	}
	ret += ")"
	return ret
}
