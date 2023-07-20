package base

import "fmt"

func main() {
	s := NewSetFrom[int](0)
	s.Set(1)
	s.Set(2)
	s.Set(4)

	fmt.Println(s)
}
