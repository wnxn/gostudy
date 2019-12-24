package binarytree

type TreeNode struct{
	V int
	Rchild *TreeNode
	Lchild *TreeNode
}

//	      1
//       /  \
//      2    3
//     / \     \
//    4   5     6
//     \       /
//      7     8
// PreOrder: 1,2,4,7,5,3,6,8
// MidOrder: 4,7,2,5,1,3,8,6
// PostOrder: 7,4,5,2,8,6,3,1
func CreateRandomTree()*TreeNode{
	head := &TreeNode{V: 1,}
	head.Lchild = &TreeNode{V:  2,}
	head.Lchild.Lchild = &TreeNode{V:4}
	head.Lchild.Lchild.Rchild = &TreeNode{V:7}
	head.Lchild.Rchild = &TreeNode{V:5}
	head.Rchild = &TreeNode{V:3}
	head.Rchild.Rchild = &TreeNode{V:6}
	head.Rchild.Rchild.Lchild = &TreeNode{V:8}
	return head
}
