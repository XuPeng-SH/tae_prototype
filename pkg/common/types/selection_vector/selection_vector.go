package selection_vector

import (
	"strconv"
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

func WithCount(count int) Option {
	return func(sv SelectionVector) SelectionVector {
		sv.InitWithCount(count)
		return sv
	}
}

func (sv *SelectionVector) InitWithCount(count int) {
	data := &SelectionData{}
	data.Data = make([]EntryT, count)
	sv.Data = data
}

func (sv *SelectionVector) InitWithData(data SelectionData) {
	sv.Data = &data
}

func (sv *SelectionVector) InitWithOther(other SelectionVector) {
	sv.Data = other.Data
}

func (sv *SelectionVector) Empty() bool {
	return (sv.Data == nil) || (len(sv.Data.Data) == 0)
}

func (sv *SelectionVector) SetIndex(index int, loc EntryT) {
	sv.Data.Data[index] = loc
}

func (sv *SelectionVector) GetIndex(index int) EntryT {
	return sv.Data.Data[index]
}

func (sv *SelectionVector) Count() int {
	if sv.Data == nil {
		return 0
	}
	return len(sv.Data.Data)
}

func (sv *SelectionVector) Slice(other SelectionVector, count int) *SelectionData {
	if count > other.Count() {
		return nil
	}
	data := SelectionData{
		Data: make([]EntryT, count),
	}
	for i := 0; i < count; i++ {
		new_idx := other.GetIndex(i)
		if (int)(new_idx) > sv.Count() {
			return nil
		}
		idx := sv.GetIndex((int)(new_idx))
		data.Data[i] = idx
	}
	return &data
}

func (sv *SelectionVector) Swap(i int, j int) {
	if i == j {
		return
	}
	tmp := sv.GetIndex(i)
	sv.SetIndex(i, sv.GetIndex(j))
	sv.SetIndex(j, tmp)
}

func (sv *SelectionVector) String(count int) string {
	ret := "SelectionVector [" + strconv.Itoa(count) + "/" + strconv.Itoa(sv.Count()) + "] ("
	if count > sv.Count() {
		count = sv.Count()
	}
	for i := 0; i < count; i++ {
		if i != 0 {
			ret += ", "
		}
		ret += strconv.Itoa((int)(sv.GetIndex(i)))
	}
	ret += ")"
	return ret
}
