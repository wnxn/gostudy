package linklist

import "fmt"


func deleteif(head *Linklist, value int)*Linklist{
	ret := &Linklist{0, head}
	prev := ret
	for head != nil{
		if head.Value==value{
			prev.Next = head.Next
			head = head.Next
		}else{
			prev = head
			head = head.Next
		}
	}
	return ret.Next
}

func deleteif2(head *Linklist, value int)*Linklist{
	if head == nil{
		return nil
	}
	for head.Value == value{
		head = head.Next
		if head==nil{
			return nil
		}
	}
	res := head
	prev := head
	for prev.Next != nil{
		if prev.Next.Value == value{
			head = head.Next
			prev.Next = head
		}else{
			prev = head
			head = head.Next
		}
	}
	return res
}

func Deleteifmain(){
	ll:=CreateRecursive([]int{2,2,1,2,2,3,2})
	Print(ll)
	fmt.Println()
	res :=deleteif2(ll,2)
	Print(res)
}
