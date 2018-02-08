package main

import (
	"fmt"
	"gostudy/retriever/mock"
	"gostudy/retriever/real"
	"time"
)

type Retriever interface{
	Get(url string)string
}

func download(r Retriever){
	fmt.Println(r.Get("http://www.imooc.com"))
}

func inspect(r Retriever){
	switch v:= r.(type){
	case mock.MockRetriever:
		fmt.Println("mock: ", v.Description)
	case *real.Retriever:
		fmt.Println("real: ", v.UserAgent)
	}
}

func main() {
	var r Retriever
	r = mock.MockRetriever{"This is a mock retriever"}
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{
		UserAgent:"Mozilla/5.0",
		TimeOut:time.Minute}
	fmt.Printf("%T %v\n", r, r)

	fmt.Println("switch case")
	inspect(r)

	fmt.Println("type assertion")
	if realRetriever,ok :=r.(*real.Retriever); ok {
		fmt.Printf("%T\n", ok)
		fmt.Println(realRetriever.UserAgent)
	}else{
		fmt.Println("Not a real retriever")
	}

//	download(r)
}
