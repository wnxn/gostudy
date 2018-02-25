package linklist

func ReverseLinklist(head *Linklist)*Linklist{
	return reverselinklist(head, nil)
}

func reverselinklist(cur, prev *Linklist)*Linklist{
	next := cur.Next
	cur.Next = prev
	if next ==nil{
		return cur
	}
	return reverselinklist(next, cur)
}

func ReverseLinklist2(head *Linklist)*Linklist{
	if head == nil || head.Next==nil{
		return head
	}
	newlink:=ReverseLinklist2(head.Next)
	head.Next.Next= head
	head.Next=nil
	return  newlink
}