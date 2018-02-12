package queue

type Queue []interface{}

// push an element to the queue
func (q *Queue)Push(i int){
	*q = append(*q,i)
}

// pop an element from a queue
func (q *Queue)Pop()int{
	ret := (*q)[0]
	*q=(*q)[1:]
	return ret.(int)
}

// check queue is empty
func (q *Queue)IsEmpty()bool{
	return len(*q) == 0
}