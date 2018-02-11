package main

import (
"fmt"
)

var mp = make([]int, 0xffff)
func maxSubarray(s string) int{

	for i,_:=range mp{
		mp[i]=-1
	}
	ret := 0
	start := 0
	arr := []rune(s)
	for i,v:=range arr{
		if mp[v] >= 0{
			start = mp[v]+1
			mp[v]= i
		}else{
			mp[v] = i
			if ret < i-start + 1{
				ret = i-start+1
			}
		}
	}
	return ret
}

func main(){
	str1 := "asdadf"
	str2:="一二三二一"
	str3 :="pwwkew"
	str4 :="bbbbbb"
	str5 :="这里是慕渴望"
	fmt.Println(maxSubarray(str1))
	fmt.Println(maxSubarray(str2))
	fmt.Println(maxSubarray(str3))
	fmt.Println(maxSubarray(str4))
	fmt.Println(maxSubarray(str5))
}