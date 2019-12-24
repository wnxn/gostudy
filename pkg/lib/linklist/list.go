package linklist

type ListNode struct{
	Val int
	Next *ListNode
}

func Get(head *ListNode)[]int{
	res := make([]int,0)
	for head != nil{
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

func Create(s []int)*ListNode{
	res := &ListNode{}
	p := res
	for _, v:=range s{
		p.Next = &ListNode{
			Val:  v,
			Next: nil,
		}
		p = p.Next
	}
	return res.Next
}
