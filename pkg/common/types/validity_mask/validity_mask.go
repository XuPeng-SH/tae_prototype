package validity_mask

import (
	"fmt"
	// log "github.com/sirupsen/logrus"
	"tae/pkg/common/types"
	"tae/pkg/common/types/constants"
)

func WhichEntry(idx types.IDX_T) types.IDX_T {
	return (idx + ((types.IDX_T)(BITS_PER_ENTRY) - 1)) / (types.IDX_T)(BITS_PER_ENTRY)
}
func GetEntryIndex(row_idx int) EntryIndex {
	ei := EntryIndex{}
	ei.Idx = row_idx / BITS_PER_ENTRY
	ei.Offset = row_idx % BITS_PER_ENTRY
	return ei
}

func (e *EntryT) IsValid(idx int) bool {
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

type Option func(ValidityMask) ValidityMask

func New(count int, options ...Option) *ValidityMask {
	vm := &ValidityMask{}
	vm.Init(count)
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

func (vm *ValidityMask) Init(count int) {
	if count < 0 {
		panic("Count should not be negtive value")
	}
	if count > constants.STANDARD_VECTOR_SIZE {
		panic(fmt.Sprintf("Too big count, should not be larger than %d", constants.STANDARD_VECTOR_SIZE))
	}

	if count == 0 {
		return
	}
	entry_count := (int)(WhichEntry((types.IDX_T)(count)))
	vm.Data = make([]EntryT, 0, entry_count)
}

func (vm *ValidityMask) Len() int {
	return len(vm.Data)
}

func (vm *ValidityMask) Reset() {
	vm.Data = vm.Data[:0]
}

func (vm *ValidityMask) GetEntry(entry_idx int) EntryT {
	if entry_idx >= vm.Len() {
		return MAX_ENTRY
	}

	return vm.Data[entry_idx]
}

func (vm *ValidityMask) SetInvalid(row_idx int) {
	ei := GetEntryIndex(row_idx)
	if ei.Idx >= vm.Len() {
		return
	}
	vm.Data[ei.Idx] &= ^(EntryT(1) << ei.Offset)
}

func (vm *ValidityMask) ValidateRows(rows int) {
	if rows >= constants.STANDARD_VECTOR_SIZE || rows < 0 {
		panic(fmt.Sprintf("Rows should be not more than %d", constants.STANDARD_VECTOR_SIZE))
	}
	if rows == 0 || vm.Len() == 0 {
		return
	}

	ei := GetEntryIndex(rows)
	for i := 0; i <= ei.Idx; i++ {
		vm.Data[i] = MAX_ENTRY
	}
}

func (vm *ValidityMask) InvalidateRows(rows int) {
	if rows >= constants.STANDARD_VECTOR_SIZE || rows < 0 {
		panic(fmt.Sprintf("Rows should be not more than %d", constants.STANDARD_VECTOR_SIZE))
	}
	if rows == 0 {
		return
	}
	if vm.Len() == 0 {
		vm.Init(rows)
	}

	ei := GetEntryIndex(rows)
	// log.Info(ei.ToString())
	for i := 0; i <= ei.Idx; i++ {
		vm.Data[i] = 0
	}
}

func (vm *ValidityMask) SetValid(row_idx int) {
	if vm.Len() == 0 {
		return
	}
	ei := GetEntryIndex(row_idx)
	if ei.Idx >= vm.Len() {
		return
	}
	vm.Data[ei.Idx] |= EntryT(1) << ei.Offset
}

func (vm *ValidityMask) IsRowValid(row_idx int) bool {
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

func (vm *ValidityMask) Slice(other ValidityMask, offset int) {
	if offset < 0 {
		panic("")
	}
	if other.AllValid() {
		vm.Reset()
		return
	}
	if offset == 0 {
		vm.Data = other.Data
		return
	}
	vm.Init(constants.STANDARD_VECTOR_SIZE)
	vm.InitAllValid()

	all_units := offset / BITS_PER_ENTRY

	if all_units != 0 {
		for idx := 0; idx+all_units < STANDARD_ENTRY_COUNT; idx++ {
			start := idx * BYTES_PER_ENTRY
			end := (idx + 1) * BYTES_PER_ENTRY
			o_start := (idx + all_units) * BYTES_PER_ENTRY
			o_end := (idx + all_units + 1) * BYTES_PER_ENTRY
			copy(vm.Data[start:end], other.Data[o_start:o_end])
		}
	}
	if sub_units := offset - all_units%BITS_PER_ENTRY; sub_units > 0 {
		for idx := 0; idx+1 < STANDARD_ENTRY_COUNT; idx++ {
		}
	}
}

func (ei *EntryIndex) ToString() string {
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
