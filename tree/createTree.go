package tree


func GetIndex(str string, ch uint8)int{
	for i,_:=range str{
		if str[i] == ch{
			return i
		}
	}
	return -1
}

func CreateTreeByPreMid(prestr, midstr string)*Node{
	if prestr == ""{
		return nil
	}
	rootChar:=prestr[0]
	root := &Node{int(rootChar),nil,nil,nil}
	rootIndex:= GetIndex(midstr, rootChar)
	root.Left = CreateTreeByPreMid(prestr[1:rootIndex + 1], midstr[:rootIndex])
	root.Right = CreateTreeByPreMid(prestr[rootIndex+1:], midstr[rootIndex+1:])
	return root
}

func CreatePostByPreMid(prestr, midstr string) string{
	if prestr== ""{
		return ""
	}
	rootChar:=prestr[0]
	rootIndex:= GetIndex(midstr, rootChar)
	lstr := CreatePostByPreMid(prestr[1:rootIndex + 1], midstr[:rootIndex])
	rstr :=  CreatePostByPreMid(prestr[rootIndex+1:], midstr[rootIndex+1:])
	return  lstr+ rstr+ string(prestr[0])
}