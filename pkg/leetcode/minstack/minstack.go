package minstack

type MinStack struct {
	normal []int
	min []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.normal = append(this.normal, x)
	if len(this.min) == 0 || this.min[len(this.min)-1] >= x{
		this.min = append(this.min, x)
	}
}


func (this *MinStack) Pop()  {
	if this.normal[len(this.normal)-1] == this.min[len(this.min)-1]{
		this.min = this.min[:len(this.min)-1]
	}
	this.normal = this.normal[:len(this.normal)-1]
}


func (this *MinStack) Top() int {
	return this.normal[len(this.normal)-1]
}


func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}