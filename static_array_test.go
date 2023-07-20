package base_test

import (
	"testing"

	"github.com/judah-caruso/base.go"
)

func TestStaticArray(t *testing.T) {
	arr := base.NewStaticArray[int](8)

	arr.Append(10)
	arr.Append(20)
	arr.Append(30)

	if arr.Len() != 3 {
		t.Errorf("array had invalid length after append, %d", arr.Len())
	}

	for i, v := range arr.Items() {
		if (i+1)*10 != v {
			t.Errorf("array had invalid value %q at index %d", v, i)
		}
	}

	val := arr.Pop()
	if val != 30 {
		t.Error("invalid value returned from pop")
	}

	arr.Remove(1)

	items := arr.Items()
	items[0] *= 2

	val = arr.Pop()
	if val != items[0] {
		t.Error("modification of items created new list")
	}
}
