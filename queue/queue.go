package queue

type Queue []interface{}

func (q *Queue)Push(i int){
	*q = append(*q,i)
}

func (q *Queue)Pop()int{
	ret := (*q)[0]
	*q=(*q)[1:]
	return ret.(int)
}

func (q *Queue)IsEmpty()bool{
	return len(*q) == 0
}