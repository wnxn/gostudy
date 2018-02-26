package main

import (
	"github.com/wnxn/gostudy/tree"
	"fmt"
)

func main() {
	root := tree.CreateMyTree1()
	for _,v:=range root.MidOrder(){
		fmt.Printf("%d, ", v)
	}
	fmt.Println()
	for root.Left != nil{
		root = root.Left
	}
	for root != nil{
		root.Print()
		root = tree.MidOrderNextNode(root)
	}
}
