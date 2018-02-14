package main

import (
	"github.com/wnxn/gostudy/tree"
	"fmt"
)

func main() {
	root:=tree.CreateTree()

	c:=root.MyTraverseWithChannel()
	maxvalue:=0
	for  value:=range c{
		if value >  maxvalue{
			maxvalue=value
		}
	}
	fmt.Printf("root max value is = %d\n", maxvalue)
}
