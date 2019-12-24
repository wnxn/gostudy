package binarytree

type SearchTreeData struct {
		left  *SearchTreeData
		data  int
		right *SearchTreeData
}

 func Bst(v int) *SearchTreeData{
 	ret := &SearchTreeData{data: v}
 	return ret
 }

 func (t *SearchTreeData) Insert(val int){
	 cur := t
	 for cur != nil{
		 if val <= cur.data{
			 if (cur.left == nil){
				 cur.left = Bst(val)
				 return
			 }
			 cur = cur.left
		 }else if val > cur.data{
			 if cur.right == nil{
				 cur.right = Bst(val)
				 return
			 }
			 cur = cur.right
		 }
	 }
	 return
 }

 func (t *SearchTreeData) MapString(f func(int) string) []string{
 	var ret []string
	 if t == nil{
		 return ret
	 }
	 for _, v:=range t.left.MapString(f){
		 ret = append(ret , v)
	 }
	 ret = append(ret, f(t.data))
	 for _, v:=range t.right.MapString(f){
		 ret = append(ret , v)
	 }
 	return ret
 }

 func (t *SearchTreeData) MapInt(f func(int) int) []int{
 	var ret []int
 	if t == nil{
 		return ret
	}
	for _, v:=range t.left.MapInt(f){
		ret = append(ret , v)
	}
	ret = append(ret, f(t.data))
	 for _, v:=range t.right.MapInt(f){
		 ret = append(ret , v)
	 }
	return ret
 }