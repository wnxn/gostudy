package main

import (
	//"unicode/utf8"
	"fmt"
	"unicode/utf8"
	"strings"
)

func main(){
	str := "Yes我爱慕课网!"
	fmt.Printf("len=%d, len(rune)=%d\n", len(str), len([]rune(str)))
	fmt.Printf("%s %X \n", []byte(str), []byte(str))

	fmt.Println("range []byte(str)")
	for _,v:=range []byte(str){
		fmt.Printf("%X ", v)
	}
	fmt.Println()

	fmt.Println("range str")
	for i,v:=range str{ // ch is a rune
		fmt.Printf("(%d,%X) ",i, v)
	}
	fmt.Println()

	// utf-8 lib
	fmt.Println("utf-8 rune count in str: ",
		utf8.RuneCountInString(str),
		)
	fmt.Println(utf8.DecodeRuneInString(str))

	// bytes array
	bytes := []byte(str)
	for len(bytes) >0{
		ch,size :=utf8.DecodeRune(bytes)
		fmt.Printf("(%c %v)", ch, size)
		bytes = bytes[size:]
	}
	fmt.Println()

	// rune array
	fmt.Println("rune array len: ", len([]rune(str)))

	str2:="rune array len: "
	fmt.Println(strings.Fields(str2))
	fmt.Println(strings.Split(str2, "a"))
	fmt.Println(strings.Index(str2, "ak"))
	fmt.Println(strings.ToUpper(str))
}
