package tree

import (
	"testing"
	"fmt"
)

func TestCreateTreeByPreMid(t *testing.T) {
	testcases := []struct {
		prestr string
		midstr string
		poststr string
	}{
		{"ABDEGCF","DBGEACF","DGEBFCA"},
		{"","",""},
		{"A","A","A"},
		{"AB","BA","BA"},
	}
	for _,test:=range testcases{
		root := CreateTreeByPreMid(test.prestr, test.midstr)
		var func1 func(*Node)string
		func1= func(node *Node)string{
			if node == nil{
				return ""
			}
			lstr:= func1(node.Left)
			rstr:=func1(node.Right)
			return fmt.Sprintf("%s%s%s",lstr,rstr,string(node.Value))
		}
		result := func1(root)
		if result != test.poststr{
			t.Errorf("prestr=%s, midstr=%s, expect poststr=%s, but actual is %s",
				test.prestr, test.midstr, test.poststr, result)
		}
	}
}

func TestCreatePostByPreMid(t *testing.T) {
	testcases := []struct {
		prestr string
		midstr string
		poststr string
	}{
		{"ABDEGCF","DBGEACF","DGEBFCA"},
		{"","",""},
		{"A","A","A"},
		{"AB","BA","BA"},
	}
	for _,test:=range testcases{
		result := CreatePostByPreMid(test.prestr, test.midstr)
		if result != test.poststr{
			t.Errorf("prestr=%s, midstr=%s, expect poststr=%s, but actual is %s",
				test.prestr, test.midstr, test.poststr, result)
		}
	}
}