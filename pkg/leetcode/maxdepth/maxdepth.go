package maxdepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var buf []*TreeNode
	var lastVisit *TreeNode
	max, cur := 0, 0
	for root != nil {
		buf = append(buf, root)
		lastVisit = root
		cur++
		if cur > max {
			max = cur
		}
		root = root.Left
	}
	for len(buf) != 0 {
		p := buf[len(buf)-1]
		cur--
		buf = buf[:len(buf)-1]
		if p.Right == nil || p.Right == lastVisit {
			lastVisit = p
		} else {
			buf = append(buf, p)
			p = p.Right
			cur++
			for p != nil {
				buf = append(buf, p)
				p = p.Left
				cur++
			}
			if cur > max {
				max = cur
			}
		}
	}
	return max
}

//func MaxDepth(root *TreeNode) int {
//	if root == nil {
//		return 0
//	}
//	var pre []*TreeNode
//	cur := root
//	maxLvl, curLvl := 0,0
//	for cur!=nil|| len(pre)!=0{
//		if cur == nil{
//			if cur == pre[len(pre)-1].Right{
//				pre = pre[:len(pre)-1]
//				curLvl--
//			}else{
//				cur = pre[len(pre)-1].Right
//			}
//		}else if cur != nil{
//			pre = append(pre, cur)
//			cur = cur.Left
//			curLvl++
//			if curLvl > maxLvl{
//				maxLvl = curLvl
//			}
//		}
//	}
//	return maxLvl-1
//}
