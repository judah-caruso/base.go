/*
# Set

A simple type that allows bitset-like functionality.
*/
package base

import (
	"math/bits"
	"unsafe"
)

// T must be a signed or unsigned integer of any size.
//
// The intended use is for flags:
//
//	type Flag uint
//
//	const (
//		FlagNone Flag = 1 << iota
//		FlagA
//		FlagB
//		FlagC
//		FlagFoo = FlagA | FlagC
//	)
//
//	flags := NewSet[Flag]()
//	flags.Set(FlagA)   // FlagA is set
//	flags.Set(FlagB)   // FlagA + FlagB is set
//	flags.Set(FlagFoo) // FlagA + FlagB + FlagC is set
type Set[T anyInt] struct {
	Value T
}

// Creates an empty Set of T.
func NewSet[T anyInt]() Set[T] {
	return Set[T]{Value: 0}
}

// Create a Set of T with the value 'v'.
func NewSetFrom[T anyInt](v T) Set[T] {
	return Set[T]{Value: v}
}

// Resets the value of a Set.
func (b *Set[T]) Reset() {
	b.Value = 0
}

// Returns the number of set bits within a Set.
func (b *Set[T]) TotalSet() int {
	size := unsafe.Sizeof(b.Value)

	switch size {
	case 1:
		return bits.OnesCount8(uint8(b.Value))
	case 2:
		return bits.OnesCount16(uint16(b.Value))
	case 4:
		return bits.OnesCount32(uint32(b.Value))
	case 8:
		return bits.OnesCount64(uint64(b.Value))
	}

	return 0
}

// Returns the number of unset bits within a Set.
func (b *Set[T]) TotalUnset() int {
	size := unsafe.Sizeof(b.Value)

	switch size {
	case 1:
		return 8 - bits.OnesCount8(uint8(b.Value))
	case 2:
		return 16 - bits.OnesCount16(uint16(b.Value))
	case 4:
		return 32 - bits.OnesCount32(uint32(b.Value))
	case 8:
		return 64 - bits.OnesCount64(uint64(b.Value))
	}

	return 0
}

// Sets a value within the Set.
func (b *Set[T]) Set(v T) {
	b.Value |= v
}

// Unsets a value within the Set.
func (b *Set[T]) Unset(v T) {
	b.Value &= ^v
}

// Checks if 'val' is within the Set.
func (b *Set[T]) Contains(val T) bool {
	return b.Value&val == val
}

// Toggles 'v' within the Set.
func (b *Set[T]) Toggle(v T) {
	if b.Contains(v) {
		b.Unset(v)
	} else {
		b.Set(v)
	}
}

// Is this really necessary?
type anyInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}
