package text

import "fmt"

type Article struct{
	S string
}

func (a Article)Read()string{
	return a.S
}

func (a Article)Write(str string){
	a.S = str
}

func (a Article)Print(){
	fmt.Printf("%T: %s\n", a, a.S)
}
