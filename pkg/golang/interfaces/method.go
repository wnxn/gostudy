package interfaces

import "fmt"

type Integer int

func (a *Integer) Add(b Integer) Integer {
	return *a + b
}

func Test() {
	var a Integer = 1

	var b Integer = 2

	var i interface{} = a

	sum := i.(Integer).Add(b)

	fmt.Println(sum)
}
