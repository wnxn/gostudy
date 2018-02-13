package tree

import (
	"fmt"
)

func (t Node)ValueReceiver(){
	fmt.Printf("treenode value=%d, type=%T\n",
		t.Value,t)
	t.Value=111
}

func(t *Node)PointerReceiver(){
	fmt.Printf("treenode value=%d, type=%T\n",
		t.Value, t)
	t.Value = 222
}

func PassValue(t Node){
	fmt.Printf("treenode value=%d, type=%T\n",
		t.Value, t)
	t.Value = 333
}

func PassPointer(t *Node){
	fmt.Printf("treenode value=%d, type=%T\n",
		t.Value, t)
	t.Value = 444
}
