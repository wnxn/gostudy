package main

import "fmt"

type people struct {
	name string
	age  int
}

func main() {
	// single int array
	sint := []int{2, 3, 5, 1, 6}
	sint[3], sint[4] = sint[4], sint[3]
	fmt.Printf("%v\n", sint)

	// struct array
	speople := []*people{
		{"wangxin", 25},
		{"liuliu", 27},
		{"gyf", 30},
	}
	speople[1], speople[0] = speople[0], speople[1]
	for _, v := range speople {
		fmt.Printf("%v,", v)
	}
	fmt.Println()

}
