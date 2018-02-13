package main

import (
	"github.com/wnxn/gostudy/tree"
	"fmt"
)


func main() {
	pnode:= new(tree.Node)
	pnode.Value=9

	node:= tree.Node{Value:6}

	fmt.Printf("pnode type is %T\n", pnode)
	fmt.Printf("node type is %T\n", node)

	fmt.Printf("Pointer receiver\n")
	pnode.PointerReceiver()
	node.PointerReceiver()

	fmt.Printf("Value receiver\n")
	pnode.ValueReceiver()
	node.ValueReceiver()

	fmt.Printf("Pass value\n")
	tree.PassValue(*pnode)
	tree.PassValue(node)

	fmt.Printf("Pass pointer\n")
	tree.PassPointer(pnode)
	tree.PassPointer(&node)

	fmt.Printf("After calling\n")
	fmt.Printf("pnode=%d\n",pnode.Value)
	fmt.Printf("node=%d\n", node.Value)


}
