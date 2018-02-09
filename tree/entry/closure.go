package main

import (
	"github.com/wnxn/gostudy/tree"
	"fmt"
)

func main() {
	p :=tree.CreateTree()
	tree.Traverse(p)
	fmt.Println()

	p.TraverseFunc(tree.ClosurePrint)
}