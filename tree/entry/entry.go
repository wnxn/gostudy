package main

import (
	"github.com/wnxn/gostudy/tree"
)

func main() {
	root :=tree.CreateTree()
	res :=tree.LowestCommonAncestorBinaryTree3(root,7,7 )
	res.Print()
}
