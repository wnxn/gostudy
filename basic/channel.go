package main

import (
	"fmt"
)


func oneroutine()string{
	ch :=make(chan string)
	go func(){ch <-"go func 1"}()
	go func(){ch <-"go func 2"}()
	go func(){ch <-"go func 3"}()
	return <-ch
}

func main() {
	fmt.Println(oneroutine())
}
