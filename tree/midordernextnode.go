package tree

import "fmt"


func MidOrderNextNode(root *Node)*Node{

	if root.Right != nil{
		root = root.Right
		for root.Left != nil{
			root = root.Left
		}
		return root
	}else{
		for root.Parent != nil && root.Parent.Right == root{
			root = root.Parent
		}
		return root.Parent
	}
}

func TraversalByMidOrderNextNode(){


	root := CreateMyTree1()
	for _,v:=range root.MidOrder(){
		fmt.Printf("%d, ", v)
	}
	fmt.Println()
	for root!= nil &&root.Left != nil{
		root = root.Left
	}
	for root != nil{
		root.Print()
		root = MidOrderNextNode(root)
	}
}