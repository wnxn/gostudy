package tree

func InvertBianryTree(root *Node)*Node{
	if root == nil{
		return nil
	}
	tmp := root.Right
	root.Right = InvertBianryTree(root.Left)
	root.Left = InvertBianryTree(tmp)
	return root
}
