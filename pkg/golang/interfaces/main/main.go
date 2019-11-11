package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

//func main() {
//	//interfaces.TestBook2()
//}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
func main() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}
