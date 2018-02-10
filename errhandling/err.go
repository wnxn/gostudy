package main

import (
	"os"
	"fmt"
	"errors"
)

func main() {
	file,err :=  os.OpenFile("fib.txt", os.O_EXCL|os.O_CREATE, 0666)
	err = errors.New("this is a custom err")
	if err != nil{
		fmt.Println(err)
		fmt.Printf("err.Error() = %s\n", err.Error())
		if v,ok:=err.(*os.PathError); ok{
			fmt.Printf("%s|%s|%s|%T|%T\n",
				v.Op,v.Path,v.Err,v,err)
		}
		return
	}
	file.Close()
}
