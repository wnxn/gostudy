package main

import (
	"os"
	"fmt"
//	"io/ioutil"
)

func readMaze(filename string)[][]int{
	file, err:=os.Open(filename)
	if err != nil{
		panic(err)
	}
	//defer file.Close()

	var row,col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	fmt.Println(row,col)

	maze := make([][]int, row)
	for i:=range maze{
		maze[i]=make([]int,col)
		//s,_:=ioutil.ReadAll(file)
		//fmt.Printf("%s",s)
		for j:=range maze[i]{
			fmt.Fscanf(file, "%d", &maze[i][j])
			fmt.Printf("(%d,%d)=%d ",i,j,maze[i][j])
		}
		fmt.Println()
	}
	return maze
}

func main() {
	maze := readMaze("maze/maze.in")
	for _, arr:=range maze{
		for _, v:=range arr{
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
