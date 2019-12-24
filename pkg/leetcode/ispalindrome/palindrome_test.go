package ispalindrome

import (
	"bytes"
	"github.com/wnxn/gostudy/pkg/lib/linklist"
	"testing"
)

func isPalindrome(s string) bool {
	length := len(s)
	s = string(bytes.ToLower([]byte(s)))
	for i, j := 0, length-1; i < j; {
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9')) {
			i++
			continue
		}
		if !((s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9')) {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		s   string
		res bool
	}{
		{
			s:   "A man, a plan, a canal: Panama",
			res: true,
		},
		{
			s:   "race a car",
			res: false,
		},
		{
			s:   "0P",
			res: false,
		},
	}
	for _, test := range tests {
		res := isPalindrome(test.s)
		if res != test.res {
			t.Errorf("IsPalindrome(%s)= %t, but expect %t", test.s, res, test.res)
		}
	}
}

//func isPalindromeList(head *linklist.ListNode)bool{
//	if head == nil || head.Next ==nil{
//		return true
//	}
//	l := &linklist.ListNode{
//		Val:  0,
//		Next: head,
//	}
//	p1 := l
//	p2 := l
//	for p1.Next != nil {
//		p1 = p1.Next
//		if p1.Next != nil {
//			p1 = p1.Next
//		}
//		p2 = p2.Next
//	}
//	p2 = p2.Next
//	var pre *linklist.ListNode
//	for p2 != nil{
//		p2.Next, p2,pre = pre, p2.Next,p2
//	}
//
//	for p1 != nil{
//		if p1.Val != head.Val{
//			return false
//		}
//		p1, head = p1.Next, head.Next
//	}
//	return true
//}

func isPalindromeList(head *linklist.ListNode)bool{
	var pre,half *linklist.ListNode
	slow, fast := head, head
	i := 1
	for fast.Next != nil{
		fast = fast.Next
		i++
		if fast.Next != nil{
			fast = fast.Next
			i++
			pre, slow, slow.Next = slow, slow.Next, pre
			half = slow.Next
		}else{
			pre, slow, slow.Next = slow, slow.Next, pre
			half = slow
		}

	}
	//if i%2==1{
	//	slow = slow.Next
	//}
	for half != nil{
		if half.Val != pre.Val{
			return false
		}
		half, pre = half.Next, pre.Next
	}
	return true
}

func TestIsPalindromeList(t *testing.T){
	tests := []struct{
		head *linklist.ListNode
		isPal bool
	}{
		{
			head:linklist.Create([]int{1,2,2,1}),
			isPal: true,
		},
		{
			head:linklist.Create([]int{1,2,1}),
			isPal: true,
		},
		{
			head:linklist.Create([]int{1,1}),
			isPal: true,
		},
		{
			head:linklist.Create([]int{1,2}),
			isPal: false,
		},
		{
			head:linklist.Create([]int{1}),
			isPal: true,
		},
	}
	for _, test := range tests{
		res := isPalindromeList(test.head)
		if res != test.isPal{
			t.Errorf("%v expect %t, but actually %t", linklist.Get(test.head) ,test.isPal,res)
		}
	}
}