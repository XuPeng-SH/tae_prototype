package validity_mask

import (
	"bytes"
	"encoding/binary"

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

func GetByteIndex(row_idx int) ByteIndex {
	bi := ByteIndex{}
	bi.Idx = row_idx / 8
	bi.Offset = row_idx % 8
	return bi
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
		arr := make([]byte, 8)
		binary.BigEndian.PutUint64(arr, uint64(MAX_ENTRY))
		for i := 0; i < cap(vm.Data); i += BYTES_PER_ENTRY {
			vm.Data = append(vm.Data, arr...)
		}
	}

	return vm
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
	entry_count := (int)(WhichEntry((types.IDX_T)(count))) * BYTES_PER_ENTRY
	vm.Data = make([]byte, 0, entry_count)
}

func (vm *ValidityMask) Len() int {
	return len(vm.Data)
}

func (vm *ValidityMask) Entries() int {
	return vm.Len() / BYTES_PER_ENTRY
}

func (vm *ValidityMask) Reset() {
	vm.Data = vm.Data[:0]
}

func (vm *ValidityMask) GetEntry(entry_idx int) EntryT {
	if entry_idx >= vm.Entries() {
		return MAX_ENTRY
	}

	var entry EntryT
	r := bytes.NewReader(vm.Data[entry_idx*BYTES_PER_ENTRY:])
	binary.Read(r, binary.BigEndian, &entry)
	return entry
}

func (vm *ValidityMask) SetInvalid(row_idx int) {
	bi := GetByteIndex(row_idx)
	if bi.Idx >= vm.Len() {
		return
	}
	vm.Data[bi.Idx] &= ^(byte(1) << bi.Offset)
}

func (vm *ValidityMask) ValidateRows(rows int) {
	if rows >= constants.STANDARD_VECTOR_SIZE || rows < 0 {
		panic(fmt.Sprintf("Rows should be not more than %d", constants.STANDARD_VECTOR_SIZE))
	}
	if rows == 0 || vm.Len() == 0 {
		return
	}

	ei := GetEntryIndex(rows)
	idx := (ei.Idx + 1) * BYTES_PER_ENTRY
	for i := 0; i < idx; i++ {
		vm.Data[i] = byte(0xff)
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
	idx := (ei.Idx + 1) * BYTES_PER_ENTRY
	for i := 0; i < idx; i++ {
		vm.Data[i] = 0
	}
}

func (vm *ValidityMask) SetValid(row_idx int) {
	if vm.Len() == 0 {
		return
	}
	bi := GetByteIndex(row_idx)
	if bi.Idx >= vm.Len() {
		return
	}
	vm.Data[bi.Idx] |= byte(1) << bi.Offset
}

func (vm *ValidityMask) IsRowValid(row_idx int) bool {
	if vm.Len() == 0 {
		return true
	}
	bi := GetByteIndex(row_idx)
	if bi.Idx >= vm.Len() {
		return true
	}

	val := vm.Data[bi.Idx] & (byte(1) << byte(bi.Offset))
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

func (ei *EntryIndex) ToString() string {
	return fmt.Sprintf("EntryIndex<%d, %d>", ei.Idx, ei.Offset)
}

func (bi *ByteIndex) ToString() string {
	return fmt.Sprintf("ByteIndex<%d, %d>", bi.Idx, bi.Offset)
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
