package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right *Node
}

func Create(val int) *Node {
	return &Node{Value: val}
}

func (t *Node) Print() {
	fmt.Println(t.Value)
}

func (t *Node) SetValue(v int) {
	if t == nil {
		return
	}
	t.Value = v
}

func Traverse(p *Node) {
	if p == nil {
		return
	}

	Traverse(p.Left)
	p.Print()
	Traverse(p.Right)
}

func CreateTree() *Node {

	p := &Node{Value: 1}           //				1
	p.Left = &Node{Value: 2}       //		2				3
	p.Right = &Node{Value: 3}      //			4       5		6
	p.Left.Right = &Node{Value: 4} //					  7
	p.Right.Left = &Node{Value: 5}
	p.Right.Right = &Node{Value: 6}
	p.Right.Left.Right = &Node{Value: 7}
	return p
}
