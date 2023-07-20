package base_test

import (
	"testing"
	"unsafe"

	"github.com/judah-caruso/base.go"
)

type Flag uint8

const (
	FlagNone Flag = 1 << iota
	FlagA
	FlagB
	FlagC
)

type FlagSet = base.Set[Flag]

func TestSet(t *testing.T) {
	set := base.NewSet[Flag]() // less preferred
	set = FlagSet{}            // equivalent, but preferred

	size := unsafe.Sizeof(set)
	if size != unsafe.Sizeof(FlagNone) {
		t.Error("set had an invalid size:", size)
	}

	bitSize := int(unsafe.Sizeof(set.Value)) * 8
	totalBits := set.TotalUnset()
	if bitSize != totalBits {
		t.Error("number of bits didn't match size of underlying type")
	}

	set.Set(FlagA)
	set.Set(FlagB)
	set.Set(FlagC)

	if !set.Contains(FlagA) {
		t.Error("expected set to contain FlagA")
	}

	if !set.Contains(FlagB) {
		t.Error("expected set to contain FlagB")
	}

	if !set.Contains(FlagC) {
		t.Error("expected set to contain FlagC")
	}

	set.Toggle(FlagA)
	if set.Contains(FlagA) {
		t.Error("set contained toggled value")
	}
	set.Toggle(FlagA)

	set.Unset(FlagB)
	if set.Contains(FlagB) {
		t.Error("set contained FlagB after removing it")
	}

	v := set.Value
	if v == 0 {
		t.Error("invalid set value")
	}

	set.Reset()

	if v == set.Value {
		t.Error("empty set had an invalid value")
	}

	if v != FlagA|FlagC {
		t.Error("invalid set value")
	}

	set = base.NewSetFrom(v)
	if set.TotalSet() != 2 {
		t.Error("invalid number of set values")
	}

	if set.TotalUnset() != 6 {
		t.Error("invalid number of unset values ")
	}
}
