package main

import (
	"fmt"
)

func maxSubarray(s string) int{
	mp := make(map[rune]int)
	ret := 0
	start := 0
	arr := []rune(s)
	for i,v:=range arr{
		if _,exist := mp[v]; exist{
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
