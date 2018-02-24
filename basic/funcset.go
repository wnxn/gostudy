package main

import "fmt"

type student struct {
	name string
	age int
}

func (s *student)setAge(a int){
	s.age =a
}

func (s student)print(){
	fmt.Printf("Name is %s. Age is %d\n", s.name,s.age)
}

func agePlus(s *student){
	s.setAge(2)
	s.print()
}

func main() {
	s1 := student{"wangx", 16}
	agePlus(&s1)
	s1.print()
}
