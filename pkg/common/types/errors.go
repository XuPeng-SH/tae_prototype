package types

import (
	"errors"
)

var (
	ErrLogicError       = errors.New("Logic Error")
	ErrCapacityOverflow = errors.New("Capacity Overflow")
)
