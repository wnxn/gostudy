package tree

<<<<<<< HEAD
import (
	"fmt"
)

func MidOrderNextNode(root *Node)*Node{
	if root == nil{
		return nil
	}
=======
import "fmt"

func MidOrderNextNode(root *Node)*Node{
	// if node.Right != nil
	//

>>>>>>> d4759401f3b1a08e7c86603ea3314d4e0766e9fe
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
<<<<<<< HEAD
	root := CreateTreeByPreMid("","")
=======
	root := CreateMyTree1()
>>>>>>> d4759401f3b1a08e7c86603ea3314d4e0766e9fe
	for _,v:=range root.MidOrder(){
		fmt.Printf("%d, ", v)
	}
	fmt.Println()
<<<<<<< HEAD
	for root!= nil &&root.Left != nil{
=======
	for root.Left != nil{
>>>>>>> d4759401f3b1a08e7c86603ea3314d4e0766e9fe
		root = root.Left
	}
	for root != nil{
		root.Print()
		root = MidOrderNextNode(root)
	}
}