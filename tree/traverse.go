package tree
func (p *Node) PreOrder() (vec []int) {
	if p == nil {
		return vec
	}
	var buf []*Node
	for p != nil || len(buf) > 0 {
		if p != nil {
			buf = append(buf, p)
			vec = append(vec, p.Value)
			p = p.Left
		} else {
			p = buf[len(buf)-1]
			buf = buf[:len(buf)-1]
			p = p.Right
		}
	}
	return vec
}

func (p *Node) MidOrder() (vec []int) {
	if p == nil {
		return vec
	}
	var buf []*Node
	for p != nil || len(buf) > 0 {
		if p != nil {
			buf = append(buf, p)

			p = p.Left
		} else {
			p = buf[len(buf)-1]
			buf = buf[:len(buf)-1]
			vec = append(vec, p.Value)
			p = p.Right
		}
	}
	return vec
}

func (p *Node) PostOrder() (vec []int) {
	if p == nil {
		return vec
	}
	var buf []*Node
	var laVisit *Node
	for p != nil {
		buf = append(buf, p)
		p = p.Left
	}
	for len(buf) != 0 {
		p = buf[len(buf)-1]
		buf = buf[:len(buf)-1]
		if p.Right == nil || p.Right == laVisit {
			vec = append(vec, p.Value)
			laVisit = p
		} else {
			buf = append(buf, p)
			p = p.Right
			for p != nil {
				buf = append(buf, p)
				p = p.Left
			}
		}
	}
	return vec
}
