package binarytree

type BinaryTreeTraverse struct{
	arr []int
}

func (b *BinaryTreeTraverse)PreOrderRecursive(head *TreeNode){
	if head == nil{
		return
	}
	b.arr = append(b.arr, head.V)
	b.PreOrderRecursive(head.Lchild)
	b.PreOrderRecursive(head.Rchild)
}

func (b *BinaryTreeTraverse)PreOrderNonRecursive(head *TreeNode){

}

func (b *BinaryTreeTraverse)GetResult()[]int{
	return b.arr
}