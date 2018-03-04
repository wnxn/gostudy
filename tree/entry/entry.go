package main

import (
	"github.com/wnxn/gostudy/tree"
)

func main() {
<<<<<<< HEAD
	tree.TraversalByMidOrderNextNode()
=======
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
>>>>>>> d4759401f3b1a08e7c86603ea3314d4e0766e9fe
}
