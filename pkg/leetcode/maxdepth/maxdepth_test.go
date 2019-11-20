package maxdepth

import "testing"

func createTree1() *TreeNode {
	res := &TreeNode{
		Val:   3,
		Left:  nil,
		Right: nil,
	}
	res.Left = &TreeNode{Val: 9}
	res.Right = &TreeNode{Val: 20}
	res.Right.Left = &TreeNode{Val: 15}
	res.Right.Right = &TreeNode{Val: 7}
	return res
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		t     *TreeNode
		depth int
	}{
		{
			t:     createTree1(),
			depth: 3,
		},
	}
	for _, test := range tests {
		res := MaxDepth(test.t)
		if res != test.depth {
			t.Errorf("MaxDepth expect %d, but actually %d", test.depth, res)
		}
	}
}
