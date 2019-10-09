package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err:= os.Open("/Users/wangxin/mygo/")
	if err != nil{
		fmt.Println(err)
	}
	defer dir.Close()
	dirs, err := dir.Readdirnames(0)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(dirs)
}
