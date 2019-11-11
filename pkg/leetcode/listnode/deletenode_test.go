package listnode

import (
	"reflect"
	"testing"
)

func createNodeList(head []int) *ListNode {
	var res, pre *ListNode
	for _, v := range head {
		cur := &ListNode{v, nil}
		if res == nil {
			res = cur
		}
		if pre != nil {
			pre.Next = cur
		}
		pre = cur
	}
	return res
}

func printNodeList(l *ListNode) []int {
	var res []int
	for l != nil {
		res = append(res, l.Val)
		l = l.Next
	}
	return res
}

func getNode(list *ListNode, val int) *ListNode {
	for list != nil {
		if list.Val == val {
			return list
		}
		list = list.Next
	}
	return list
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		head   []int
		node   int
		output []int
	}{
		{
			head:   []int{4, 5, 1, 9},
			node:   5,
			output: []int{4, 1, 9},
		},
		{
			head:   []int{4, 5, 1, 9},
			node:   1,
			output: []int{4, 5, 9},
		},
	}
	for _, test := range tests {
		link := createNodeList(test.head)
		node := getNode(link, test.node)
		DeleteNode(node)
		if !reflect.DeepEqual(printNodeList(link), test.output) {
			t.Errorf("DeleteNode = %d, but expect %d",
				printNodeList(link), test.output)
		}
	}
}
