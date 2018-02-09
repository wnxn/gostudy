package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"bufio"
	"reflect"
	"runtime"
	"math"
)

func main() {
	var a = 142 //1000 1110(142)
	var b = 23 //10111(23)
	var file = "abc.txt"
	ifusage()
	fmt.Println(grade(59),
		grade(78),
		grade(99),
		grade(100))
	fmt.Println(inttobin(a))
	fmt.Println(inttobin(b))
	printfile(file)
	if result,err:=eval("/",21,9); err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result)
	}
	fmt.Println(sum(1,2,3,4,5))
	fmt.Println(apply(func (a float64, b float64) float64{
		return math.Pow(float64(a),float64(b))
	}, 3,2))

	a,b =swap(a,b)
	fmt.Println(a,b)

}

func ifusage() {
	const filename = "abc.txt"

	if contents, err := ioutil.ReadFile(filename); err == nil {
		fmt.Println( contents)
	} else {
		fmt.Println(err)
		return
	}
	//fmt.Println(contents)
}

func grade(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("input socre %d err", score))
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func inttobin(i int) string{
	var ret string = ""
	for ;i>0;i/=2{
		ret = strconv.Itoa(i%2) + ret
	}
	return ret
}

func printfile(filename string){
	fp, err := os.Open(filename)
	if(err != nil){
		panic(err)
	}

	scanner := bufio.NewScanner(fp);
	for scanner.Scan() {
		fmt.Println(scanner.Text());
	}

}

func forever(){
	for{
		fmt.Println("abc")
	}
}

func eval(op string, a int, b int) (int, error){
	switch op{
	case "+":
		return a+b,nil
	case "-":
		return a-b,nil
	case "*":
		return a*b,nil
	case "/":
		ret,_ :=div(a,b)
		return ret,nil
	default:
		return 0,fmt.Errorf("err ops %s\n", op)
	}
}

func div(a, b int) (q, r int){
	q = a/b
	r = a%b
	return q,r
}

func sum(a...int) (ret int){
	for _,i:=range a{
		ret += i
	}
	return ret
}

func apply(op func (float64, float64) float64, a , b float64)(ret float64){
	p := reflect.ValueOf(op).Pointer()
	opname := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling func %s " + "with args %f, %f\n", opname, a,b)
	return op(a,b)
}

func pow(a int, b int) int{
	return int(math.Pow(float64(a), float64(b)))
}

func swap(a,b int) (int,int){
	return b,a
}