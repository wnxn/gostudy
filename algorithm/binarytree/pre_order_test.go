package binarytree

import (
	"reflect"
	"testing"
)

func TestPreOrderRecursive(t *testing.T) {
	tests := []struct{
		name string
		tree *TreeNode
		arr []int
	}{
		{
			name: "randonTree",
			tree: CreateRandomTree(),
			arr: []int{1,2,4,7,5,3,6,8},
		},
	}
	for _, test := range tests{
		obj := BinaryTreeTraverse{}
		obj.PreOrderRecursive(test.tree)
		if !reflect.DeepEqual(test.arr, obj.GetResult()){
			t.Errorf("PreOrderRecursive(%s) expect %v, but actually %v", test.name, test.arr, obj.GetResult())
		}
	}
}
