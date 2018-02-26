package tree

import (
	"fmt"
)

type Node struct {
	Value       int
	Left, Right,Parent *Node
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

func (t *Node)SetLeft(left *Node){
	t.Left = left
	if left != nil{
		left.Parent = t
	}
}

func (t *Node)SetRight(right *Node){
	t.Right = right
	if right != nil{
		right.Parent = t
	}
}

func (t *Node)GetLeft()*Node{
	return t.Left
}

func (t *Node)GetRight()*Node{
	return t.Right
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

func CreateMyTree1() *Node {
	//				1
	//		2				3
	//			4       5		6
	//					  7
	// preorder: 1243576
	// midorder: 2415736
	// postorder: 4275631
	p := &Node{1,nil,nil,nil}
	p.SetLeft(&Node{Value: 2})
	p.SetRight( &Node{Value: 3} )
	p.GetLeft().SetRight(&Node{Value: 4})
	p.GetRight().SetLeft(&Node{Value: 5})
	p.GetRight().SetRight(&Node{Value: 6})
	p.GetRight().GetLeft().SetRight(&Node{Value: 7})
	return p
}

func CreateMyTree2() *Node {
	//				1
	//		2				3
	//	4		5       		6
	//		  7
	// preorder: 1243576
	// midorder: 2415736
	// postorder: 4275631
	p := &Node{Value: 1}
	p.SetLeft(&Node{Value: 2})
	p.SetRight( &Node{Value: 3} )
	p.GetLeft().SetLeft(&Node{Value:4})
	p.GetLeft().SetRight(&Node{Value:5})
	p.GetRight().SetRight(&Node{Value: 6})
	p.GetLeft().GetRight().SetLeft(&Node{Value: 7})
	return p
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}