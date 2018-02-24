package oop

import "fmt"

type people struct {
	age int
	name string
}

func (p *people)setName(s string){
	p.name =s
}

func (p *people)setAge(a int){
	p.age = a
}

func (p people)print(){
	fmt.Printf("%s is %d\n",p.name,p.age)
}

