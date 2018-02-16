package main

import (
	"fmt"
)

var maze=[][]int{
	{999,-1,0,0,0},
	{0,0,0,-1,0},
	{0,-1,0,-1,0},
	{-1,-1,-1,0,0},
	{0,-1,0,0,-1},
	{0,-1,0,0,0},
}

type mazeLoc struct{
	row int
	col int
}

var row=len(maze)
var col=len(maze[0])

func oneStep(queue []mazeLoc){
	cur := (queue)[0]
	queue= (queue)[1:]
	fmt.Printf("cur is (%d,%d)\n", cur.row,cur.col)
	for i:=0;i<4;i++{
		tmp := cur
		switch i{
		case 0:
			tmp.row--
		case 1:
			tmp.col--
		case 2:
			tmp.row++
		case 3:
			tmp.col++
		}
		if tmp.row<0 || tmp.row>=row || tmp.col<0 || tmp.col>=col || maze[tmp.row][tmp.col] !=0{
			fmt.Printf("(%d,%d) continues\n", tmp.row, tmp.col)
			continue
		}else{
			fmt.Printf("(%d,%d) not continues\n", tmp.row, tmp.col)
			maze[tmp.row][tmp.col] = maze[cur.row][cur.col] + 1
			queue=append(queue, tmp)
		}
		if tmp.row == row-1 && tmp.col == col-1{
			return
		}
	}
	oneStep(queue)
}

func main() {
	q :=[]mazeLoc{
		{0,0},
	}
	oneStep(q)
	fmt.Printf("path len = (%d)\n", maze[row-1][col-1]-999)
}
