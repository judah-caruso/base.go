/*
# Static Array

A simple fixed-size array. Not super useful yet.
*/
package base

type StaticArray[T any] struct {
	items []T
	count int
}

func NewStaticArray[T any](cap uint) StaticArray[T] {
	return StaticArray[T]{
		items: make([]T, cap),
		count: 0,
	}
}

// Slice of currently occupied items.
//
// Allows for easier iteration:
//
//	arr := NewStaticArray[int](10)
//	// ...
//	for i, v := range arr.Items() {
//		// ...
//	}
func (a *StaticArray[T]) Items() []T {
	return a.items[:a.count]
}

// Removes and returns the last element of the array.
func (a *StaticArray[T]) Pop() T {
	item := a.items[a.count-1]
	a.count -= 1
	return item
}

// Appends 'v' to the end of the array. Returns a pointer to the stored value.
func (a *StaticArray[T]) Append(v T) *T {
	a.items[a.count] = v
	a.count += 1
	return &a.items[a.count-1]
}

// Ordered removal of the value at index 'idx'.
func (a *StaticArray[T]) Remove(idx int) {
	var zero T
	for i := idx; i < a.count-2; i += 1 {
		a.items[i] = a.items[i+1]
		a.items[i+1] = zero
	}

	a.count -= 1
}

// Unordered removal of the value at index 'idx'.
func (a *StaticArray[T]) UnorderedRemove(idx int) {
	if idx > 1 && idx != a.count-1 {
		var zero T
		a.items[idx] = a.items[a.count-1]
		a.items[a.count-1] = zero
	}

	a.count -= 1
}

// Sets all values within the array to their zero value and resets the count.
func (a *StaticArray[T]) Reset() {
	var zero T
	for i := range a.items {
		a.items[i] = zero
	}

	a.count = 0
}

// Number of occupied items within the array.
func (a *StaticArray[T]) Len() int {
	return a.count
}

// Capacity of the array.
func (a *StaticArray[T]) Cap() int {
	return cap(a.items)
}
