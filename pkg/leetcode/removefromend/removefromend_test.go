package removefromend

import (
	"reflect"
	"testing"
)

type ListNode struct {
	     Val int
	     Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	l := &ListNode{
		Val:  0,
		Next: head,
	}
	p1 := l
	p2 := l
	if head.Next == nil && n == 1{
		return nil
	}
	for i:=0;p1.Next != nil;i++{
		p1 = p1.Next
		if i>=n{
			p2 = p2.Next
		}
	}
	p2.Next = p2.Next.Next
	return l.Next

}

func CreateLinkList(i []int)*ListNode{
	res := &ListNode{}
	c := res
	for _,v:= range i{
		c.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		c= c.Next
	}
	return res.Next
}

func GetLinkList(head *ListNode)[]int{
	res := make([]int,0,0)
	for head != nil{
		res = append(res,head.Val)
		head = head.Next
	}
	return res
}

func TestRemoveFromEnd(t *testing.T){
	tests := []struct{
		head []int
		n int
		res []int
	}{
		//{
		//	head: []int{1,2},
		//	n: 1,
		//	res: []int{1},
		//},
		//{
		//	head: []int{1,2,3},
		//	n: 2,
		//	res:[]int{1,3},
		//},
		//{
		//	head: []int{1},
		//	n:1,
		//	res:[]int{},
		//},
		{
			head: []int{1,2},
			n: 2,
			res:[]int{2},
		},
	}
	for _, test :=range tests{
		res := RemoveNthFromEnd(CreateLinkList(test.head), test.n)
		ressliec := GetLinkList(res)
		if !reflect.DeepEqual(ressliec, test.res){
			t.Errorf("expect %v, but actually %v", test.res, ressliec)
		}
	}
}
