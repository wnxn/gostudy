package tree

import (
	"math"
)

func LowestCommonAncestorBinaryTree1(root *Node, p int, q int)*Node{
	if root==nil || root.Value ==p || root.Value==q{
		return root
	}
	left := LowestCommonAncestorBinaryTree1(root.Left, p,q)
	right:=LowestCommonAncestorBinaryTree1(root.Right,p,q)
	if left ==nil{
		return right
	}else{
		if right == nil{
			return left
		}else{
			return root
		}
	}
}

func LowestCommonAncestorBinaryTree2(root *Node, p int, q int)*Node{
	if root == nil || root.Value == p || root.Value == q{
		return root
	}
	var pathp, pathq,temp []*Node
	temp = append(temp,root)
	var prev *Node=nil
	for len(pathp)==0 ||len(pathq)==0  {
		root =temp[len(temp)-1]
		if root.Value == p{
			for _, v:=range temp{
				pathp = append(pathp, v)
			}
		}
		if root.Value==q{
			for _, v:=range temp{
				pathq=append(pathq,v)
			}
		}

		if (root.Left==nil && root.Right==nil) ||
			(root.Right ==nil && prev ==root.Left) ||
			(root.Right != nil && prev==root.Right){
			temp = temp[:len(temp)-1]
			prev = root
		}else{
			if root.Left == nil || prev == root.Left{
				temp = append(temp, root.Right)
			}else{
				temp = append(temp, root.Left)
			}
		}
	}
	n:=math.Min(float64(len(pathp)),float64(len(pathq)))
	for i:=1;i<int(n);i++ {
		if pathp[i] != pathq[i] {
			return pathp[i-1]
		}
	}
	return pathp[int(n)-1]
}

func LowestCommonAncestorBinaryTree3(root *Node, p int, q int)*Node{
	if root == nil || root.Value ==p || root.Value == q{
		return root
	}
	var pathp,pathq,temp []*Node
	temp = append(temp, root)
	var prev *Node = nil
	for len(pathp) ==0|| len(pathq)==0{
		root = temp[len(temp)-1]
		if root.Value == p{
			for _, v :=range temp{
				pathp = append(pathp, v)
			}
		}
		if root.Value == q{
			for _,v:=range temp{
				pathq=append(pathq,v)
			}
		}
		if root.Left == nil && root.Right==nil ||
			root.Right == nil && prev == root.Left ||
			root.Right != nil && prev == root.Right{
			temp = temp[:len(temp)-1]
			prev = root
		}else{
			if root.Left == nil || root.Left == prev{
				temp = append(temp, root.Right)
			}else{
				temp = append(temp, root.Left)
			}
		}
	}
	n := int(math.Min(float64(len(pathp)), float64(len(pathq))))
	for i:=0;i<n;i++{
		if pathp[i] != pathq[i] {
			return pathp[i-1]
		}
	}
	return pathp[n-1]
}
