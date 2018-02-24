package oop


func Myset(){
	wx:=people{6,"wangx"}
	lujq:=people{
		age: 30,
		name: "lujq",
	}
	chenyj:=new(people)
	wx.age=21
	wx.print()
	lujq.print()
	chenyj.print()
}
