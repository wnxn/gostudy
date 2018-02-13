package main

import (
	"fmt"
	"github.com/wnxn/gostudy/tree"
)

func main() {
	p3 := new(tree.Node)
	p3.Value=3
	p3.Left = &tree.Node{Value:3,Left:nil,Right:nil}
	fmt.Println("Print2 problem")
	fmt.Printf("%T\n", p3)
	p3.Print()
}
