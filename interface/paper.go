package text

import "fmt"

type Paper struct{
	S string
}

func (p Paper)Read()string{
	return p.S
}

func (p Paper)Print(){
	fmt.Printf("%T: %s\n",p,p.S)
}
