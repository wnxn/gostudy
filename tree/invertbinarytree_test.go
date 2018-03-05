package tree

import "testing"

var testcases = []struct {
	original *Node
	preorder []int
	midorder []int
}{{CreateMyTree1(),CreateMyTree2().PreOrder(), CreateMyTree2().MidOrder()}}

func TestInvertBianryTree(t *testing.T) {
	for _,test:=range testcases{
		res := InvertBianryTree(test.original)
		pre := res.PreOrder()
		mid :=res.MidOrder()
		for i,_:=range test.preorder{
			if test.preorder[i] != pre[i] || test.midorder[i] != mid[i]{
				t.Errorf("TestInvertBinaryTree error expect %v, but result is %v", test.preorder, pre)
			}
		}
	}
}
