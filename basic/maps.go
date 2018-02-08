package main

import "fmt"

func printMap(m map[string]int){
	fmt.Println(m, len(m))
}

type Vertex struct{
	Lat float64
	Long float64
}

func MapOfStruct(){
	var m map[string]Vertex=map[string]Vertex{}
	fmt.Printf("%T\n", m)
	a := Vertex{Lat:40.68433, Long:-74.39967,}
	m ["Shanghai"]=a
	m["Bei Jing"] = Vertex{
		Lat:39.26, Long:115.25,
	}
	m["Guang Dong"] = Vertex{
		Lat:39.26, Long:115.25,
	}
	fmt.Println(m)
	for lat,k:=range m{
		fmt.Println(lat, k)
	}
	delete(m,"Bei Jing")
	fmt.Println(m)

	mp:=make(map[string]int)
	mp["SH"] = 021
	mp["BJ"] = 010
	fmt.Println(mp)

	var m1 map[string]Vertex
	m2 := map[string]Vertex{}
	m3 := make(map[string]Vertex)
	fmt.Printf("m1: %T,len=%d, nil=%q\n", m1, len(m1), m1==nil)
	fmt.Printf("m2: %T,len=%d, nil=%q\n", m2, len(m2), m2==nil)
	fmt.Printf("m13: %T,len=%d, nil=%q\n", m3, len(m3), m3==nil)
}

func main(){
	/*
	var m1 map[string]int
	m2 := make(map[string]int)
	m3 := map[string]int{
		"wangx":25,
		"lujq":34,
		"chenyj":37,
		}
	printMap(m1)
	printMap(m2)
	printMap(m3)

	fmt.Println("traverse map m3:")
	for k,v:=range m3{
		fmt.Println(k,v)
	}

	if age,ok:=m3["wangx"]; ok{
		fmt.Println(age)
	}else{
		fmt.Println("key not found")
	}

	fmt.Println("delete key")
	delete(m3, "wangx")
	printMap(m3)
*/
	MapOfStruct()
}
