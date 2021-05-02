package vmask

import (
	"fmt"
	// log "github.com/sirupsen/logrus"
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
)

func WhichEntry(idx types.IDX_T) types.IDX_T {
	return (idx + ((types.IDX_T)(BITS_PER_ENTRY) - 1)) / (types.IDX_T)(BITS_PER_ENTRY)
}
func GetEntryIndex(row_idx types.IDX_T) EntryIndex {
	ei := EntryIndex{}
	ei.Idx = row_idx / BITS_PER_ENTRY
	ei.Offset = row_idx % BITS_PER_ENTRY
	return ei
}

func (e *EntryT) IsValid(idx types.IDX_T) bool {
	if idx < 0 || idx >= BITS_PER_ENTRY {
		panic(fmt.Sprintf("Invalid idx %d", idx))
	}
	val := *e & (EntryT(1) << idx)
	return val > 0
}

func (e *EntryT) AllValid() bool {
	return *e == MAX_ENTRY
}

func (e *EntryT) NoneValid() bool {
	return *e == 0
}

func (e *EntryT) String() string {
	ret := "Entry ("
	if *e == 0 {
		ret += "--"
	} else if *e == MAX_ENTRY {
		ret += "++"
	} else {
		for i := BITS_PER_ENTRY; i >= 1; i-- {
			// log.Info(i)
			if *e&(EntryT(1)<<i) > 0 {
				ret += "."
			} else {
				ret += "X"
			}
		}
	}
	ret += ")"
	return ret
}

var (
	_ IValidityMask = (*ValidityMask)(nil)
)

type Option func(ValidityMask) ValidityMask

func New(count types.IDX_T, options ...Option) IValidityMask {
	vm := &ValidityMask{}
	vm.MakeRoom(count)
	for _, option := range options {
		*vm = option(*vm)
	}
	if len(vm.Data) == 0 && cap(vm.Data) != 0 {
		vm.InitAllValid()
	}

	return vm
}

func (vm *ValidityMask) InitAllValid() {
	for i := 0; i < cap(vm.Data); i++ {
		vm.Data = append(vm.Data, MAX_ENTRY)
	}
}

func (vm *ValidityMask) MakeRoom(count types.IDX_T) {
	if count > constants.STANDARD_VECTOR_SIZE {
		panic(fmt.Sprintf("Too big count, should not be larger than %d", constants.STANDARD_VECTOR_SIZE))
	}

	if count == 0 {
		return
	}
	if cap(vm.Data) > 0 {
		return
	}
	entry_count := (int)(WhichEntry((types.IDX_T)(count)))
	vm.Data = make([]EntryT, 0, entry_count)
}

func (vm *ValidityMask) Len() types.IDX_T {
	return (types.IDX_T)(len(vm.Data))
}

func (vm *ValidityMask) Reset() {
	vm.Data = vm.Data[:0]
}

func (vm *ValidityMask) GetEntry(entry_idx types.IDX_T) EntryT {
	if entry_idx >= vm.Len() {
		return MAX_ENTRY
	}

	return vm.Data[entry_idx]
}

func (vm *ValidityMask) SetInvalid(row_idx types.IDX_T) {
	ei := GetEntryIndex(row_idx)
	if ei.Idx >= vm.Len() {
		return
	}
	vm.Data[ei.Idx] &= ^(EntryT(1) << ei.Offset)
}

func (vm *ValidityMask) ValidateRows(rows types.IDX_T) {
	if rows >= constants.STANDARD_VECTOR_SIZE {
		panic(fmt.Sprintf("Rows should be not more than %d", constants.STANDARD_VECTOR_SIZE))
	}
	if rows == 0 || vm.Len() == 0 {
		return
	}

	ei := GetEntryIndex(rows)
	for i := types.IDX_0; i <= ei.Idx; i++ {
		vm.Data[i] = MAX_ENTRY
	}
}

func (vm *ValidityMask) InvalidateRows(rows types.IDX_T) {
	if rows >= constants.STANDARD_VECTOR_SIZE || rows < 0 {
		panic(fmt.Sprintf("Rows should be not more than %d", constants.STANDARD_VECTOR_SIZE))
	}
	if rows == 0 {
		return
	}
	if vm.Len() == 0 {
		vm.MakeRoom(rows)
	}

	ei := GetEntryIndex(rows)
	// log.Info(ei.String())
	for i := types.IDX_0; i <= ei.Idx; i++ {
		vm.Data[i] = 0
	}
}

func (vm *ValidityMask) SetValid(row_idx types.IDX_T) {
	if vm.Len() == 0 {
		return
	}
	ei := GetEntryIndex(row_idx)
	if ei.Idx >= vm.Len() {
		return
	}
	vm.Data[ei.Idx] |= EntryT(1) << ei.Offset
}

func (vm *ValidityMask) IsRowValid(row_idx types.IDX_T) bool {
	if vm.Len() == 0 {
		return true
	}
	ei := GetEntryIndex(row_idx)
	if ei.Idx >= vm.Len() {
		return true
	}

	val := vm.Data[ei.Idx] & (EntryT(1) << ei.Offset)
	return val > 0
}

func (vm *ValidityMask) AllValid() bool {
	if vm.Len() == 0 {
		return true
	}
	for _, b := range vm.Data {
		if b != 0 {
			return false
		}
	}
	return true
}

func (vm *ValidityMask) GetData() []EntryT {
	return vm.Data
}

func (vm *ValidityMask) Slice(other IValidityMask, offset types.IDX_T) {
	if other.AllValid() {
		vm.Reset()
		return
	}
	if offset == 0 {
		vm.Data = other.GetData()
		return
	}
	vm.MakeRoom(constants.STANDARD_VECTOR_SIZE)
	vm.InitAllValid()

	all_units := offset / BITS_PER_ENTRY
	data := other.GetData()

	if all_units != 0 {
		for idx := types.IDX_0; idx+all_units < STANDARD_ENTRY_COUNT; idx++ {
			vm.Data[idx] = data[idx+all_units]
		}
	}
	if sub_units := offset - all_units%BITS_PER_ENTRY; sub_units > 0 {
		idx := types.IDX_0
		for ; idx+1 < STANDARD_ENTRY_COUNT; idx++ {
			vm.Data[idx] = (data[idx] >> sub_units) | (data[idx+1] << (BITS_PER_ENTRY - sub_units))
		}
		vm.Data[idx] >>= sub_units
	}
}

func (vm *ValidityMask) Combine(other IValidityMask, count types.IDX_T) {
	if other.AllValid() {
		return
	}
	if vm.AllValid() {
		vm.Data = other.GetData()
		return
	}
	old_data := vm.Data
	vm.MakeRoom(constants.STANDARD_VECTOR_SIZE)
	vm.InitAllValid()

	ei := GetEntryIndex(count)
	for i := types.IDX_0; i <= ei.Idx; i++ {
		vm.Data[i] = old_data[i] & other.GetData()[i]
	}
}

func (vm *ValidityMask) String(count types.IDX_T) string {
	ret := fmt.Sprintf("ValidityMask (%v)[", count)
	for i := types.IDX_0; i < count; i++ {
		if vm.IsRowValid(i) {
			ret += "."
		} else {
			ret += "X"
		}
	}

	ret += "]"
	return ret
}

func (ei *EntryIndex) String() string {
	return fmt.Sprintf("EntryIndex<%d, %d>", ei.Idx, ei.Offset)
}

func WithOriginal(original ValidityMask) Option {
	return func(vm ValidityMask) ValidityMask {
		if vm.Len() == 0 {
			panic("Count should be specified before")
		}
		copy(vm.Data, original.Data)
		return vm
	}
}
