package main

import (
	"github.com/wnxn/gostudy/interface"
	"fmt"
)

func readonly(obj text.RO){
	ptr, ok:= obj.(text.Paper)
	if ok == false{
		fmt.Printf("false: %T\n",ptr)
		return
	}else{
		fmt.Printf("true: %T\n", ptr)
	}
	fmt.Printf("readonly: %T: %s\n", obj, obj.Read())
}

func readwrite(obj text.RW){

	obj.Write("readwrite "+ obj.Read())
	fmt.Printf("readwrite: %T: %s\n", obj, obj.Read())
}

func assertSwitch(obj text.RO){
	switch obj.(type){
	case text.Paper:
		fmt.Printf("type=%T, read=%s\n", obj, obj.Read())
	case text.Article:
		fmt.Printf("type=%T, read=%s\n", obj, obj.Read())
	}
}

func assertArticle(obj text.RO){
	if v,ok:= obj.(text.Article);ok{
		fmt.Printf("obj is a Article type: type=%T, read=%s\n", v, v.Read())
	}else{
		fmt.Printf("obj not a Article type: type=%T, read=%s\n", v, v.Read())
	}
}


func main() {
	var ro text.RO
	var rw text.RW
	ro =text.Paper{"This is a long paper"}
	rw =text.Article{"This is a short article"}

	fmt.Println(">Assert Switch")
	assertSwitch(ro)
	assertSwitch(rw)

	fmt.Println(">Assert Article")
	assertArticle(ro)
	assertArticle(rw)
	v,ok:=ro.(text.Article)
	fmt.Printf("obj (%s) a Article type: type=%T, read=%s\n",ok, v, v.Read())
}
