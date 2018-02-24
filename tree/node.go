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

func (t Node) Print() {
	fmt.Printf("%d ",t.Value)
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

func (p *Node)MyTraverseFunc(f func (*Node)){
	if p == nil {
		return
	}
	p.Left.TraverseFunc(f)
	f(p)
	p.Right.TraverseFunc(f)
}

func (p *Node)MyTraverseWithChannel()chan int{
	ret := make(chan int)
	go func(){
		p.MyTraverseFunc(func (node *Node){
			ret <- node.Value
		})
		close(ret)
	}()
	return ret
}

func ClosurePrint(p *Node){
	fmt.Printf("(%T,%d) ", p, p.Value)
}

func printTreenodeSlice(slice []*Node){
	for _, v:=range slice{
		fmt.Printf("%d, ", v.Value)
	}
	fmt.Println()
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

func CreateNode(value int) *Node {
	return &Node{Value: value}
}