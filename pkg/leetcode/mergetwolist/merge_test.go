package mergetwolist

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val   int
	Next *ListNode
}

func createList(s []int)*ListNode{
	res := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := res
	for i:=range s{
		cur.Next = &ListNode{
			Val:  s[i],
			Next: nil,
		}
		cur = cur.Next
	}
	return res.Next
}

func getList(l *ListNode)[]int{
	res := make([]int,0)
	for l!=nil{
		res=append(res, l.Val)
		l=l.Next
	}
	return res
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	for l1 != nil && l2 != nil{
		if l1.Val < l2.Val{
			cur.Next = l1
			l1 = l1.Next
		}else{
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil{
		cur.Next = l1
	}
	if l2 != nil{
		cur.Next = l2
	}
	return res.Next
}

func TestMergeTwoList(t *testing.T){
	tests := []struct{
		l1 *ListNode
		l2 *ListNode
		l3 *ListNode
	}{
		{
			l1: createList([]int{1,2,4}),
			l2: createList([]int{1,3,4}),
			l3: createList([]int{1,1,2,3,4,4}),
		},
	}
	for _, test := range tests{
		t.Logf("%v",getList(test.l3))
		res := mergeTwoLists(test.l1, test.l2)
		if !reflect.DeepEqual(getList(res), getList(test.l3)){
			t.Errorf("expect %v, but actually %v",  getList(test.l3),getList(res))
		}
	}
}