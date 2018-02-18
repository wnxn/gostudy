package main

import (
	"regexp"
	"fmt"
)

const text=`
email1 ccmouse@gmail.com
email2 is abc@qq.com
email3 is kk123@126.com.cn
`

func main() {
	re,err := regexp.Compile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	if err != nil{
		panic(err)
	}
	result:=re.FindAllStringSubmatch(text,-1)
	for _,i:=range result{
		for _,j:=range i{
			fmt.Print(j)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
