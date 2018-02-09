package main

import (
	"fmt"
	"github.com/wnxn/gostudy/tree"
)

func main() {
	var p1 tree.Node
	//			3
	//		9 		0
	//		  99  88
	p2 := tree.Node{Value: 3}
	arr := []tree.Node{
		{Value: 3},
		p2,
	}
	p2.Left = &tree.Node{Value: 9, Left: nil, Right: nil}
	p2.Right = new(tree.Node)
	p2.Left.Right = tree.Create(99)
	p2.Right.Left = tree.Create(88)
	fmt.Println(p1, p2, arr, p2.Left, p2.Right, p2.Left.Right)
	fmt.Println("treeNode method")
	p2.Left.Print()
	p2.Left.SetValue(8)
	p2.Left.Print()

	// call nil *treeNode function
	var p3 *tree.Node
	if(p3 == nil){
		fmt.Println("p3 is nil")
		p3.SetValue(9)
	}

	tree.Traverse(&p2)
	// preOrder
	p4 := tree.CreateTree()
	fmt.Println("preOrder:")
	fmt.Println(p4.PreOrder())
	// midOrder
	fmt.Println("midOrder:")
	fmt.Println(p4.MidOrder())
	// postOrder
	fmt.Println("postOrder:")
	fmt.Println((*p4).PostOrder())

}
