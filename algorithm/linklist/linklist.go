package linklist

import (
	"fmt"
)

type Linklist struct {
	Value int
	Next *Linklist
}

func Creat(len int)*Linklist{
	var head, cur, prev *Linklist
	for i:=0;i<len;i++{
		cur = &Linklist{i,nil}
		if i == 0{
			head = cur
			prev = head
		}else{
			prev.Next = cur
			prev = cur
		}
	}
	return head
}

func CreateRecursive(data []int)*Linklist{
	if len(data) <1{
		return nil
	}
	cur := &Linklist{
		Value:data[0],
		Next:nil,
		}
	cur.Next = CreateRecursive(data[1:])
	return cur
}

func Print(head *Linklist){
	for head != nil{
		fmt.Printf("%d,", head.Value)
		head = head.Next
	}
}