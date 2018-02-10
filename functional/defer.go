package main

import (
	"github.com/wnxn/gostudy/functional/fib"
	"fmt"
	"os"

	"bufio"
	"io"
)

func tryDefer(){
	for i:=0; i< 100;i++{
		defer fmt.Println(i)
	}
}

func main() {
	fmt.Println(">>fibonacci")
	f := fib.Fibonacci()
	file,err := os.Create("fib.txt")
	if err!=nil{
		panic("open file err")
	}
	defer file.Close()
	if _,ok:=interface{}(file).(io.Writer); ok{
		fmt.Println("is a io.Writer type")
	}else{
		panic("not a io.Writer type")
	}
	writer :=bufio.NewWriter(file)
	defer writer.Flush()
	for i:=1;i<10;i++{
		fmt.Fprintln(writer, f())
	}

	fmt.Println(">> defer arguements")
	tryDefer()
}
