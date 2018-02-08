package ch02

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 66
	bb = true
	cc = "wer"
)

func main() {
	fmt.Println("hello world")
	varDefine()
	varDeduction()
	fmt.Println(aa, bb, cc)
	triangle()
	euler()
	consts()
	enum()
}

func varDefine() {
	var a = 22
	var str = "wangx"
	fmt.Printf("%d %q\n", a, str)
}

func varDeduction() {
	a := 22
	b := false
	str := "wangx"
	fmt.Printf("%d %q %t\n", a, str, b)
}

func triangle() {
	var a = 3 + 4i
	fmt.Println(cmplx.Abs(a))
	var x, y int = 3, 5
	var c = int(math.Sqrt(float64(x*x + y*y)))
	fmt.Println(c)
}

func euler() {
	fmt.Printf("%.3f\n", cmplx.Exp(math.Pi*(1i))+1)
}

func consts() {
	const a = 3
	const b = 4
	const c = "wadwd"
	const d = true
	var x = math.Sqrt(a*a + b*b)
	fmt.Println(x)
	fmt.Printf("%s, %t\n", c, d)
	const f float64 = 211
	fmt.Println((f - 32) * 5 / 3)
}

func enum() {
	const (
		b = 1 << (10 * iota)
		_
		mb
		gb
		tb
		_
		eb
		zb
		yb
	)
	fmt.Println(b, mb, gb, tb, eb, yb/zb)
}
