package main

import (
	"os"
	"fmt"
)

func readMaze(filename string)[][]int{
	file, err:=os.Open(filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	var row,col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i:=range maze{
		maze[i]=make([]int,col)
		for j:=range maze[i]{
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze
}

func printMaze(maze [][]int){
	for _, arr:=range maze{
		for _, v:=range arr{
			fmt.Printf("%3d ", v)
		}
		fmt.Println()
	}
}

type point struct{
	row,col int
}

var directions=[]point{
	{-1,0},{0,-1},{1,0},{0,1},
}

func (p point)add(q point)point{
	return point{p.row+q.row,p.col+q.col}
}

func getValue(m [][]int, loc point) (int,bool){
	row := len(m)
	if row == 0{
		return 0,false
	}
	col := len(m[0])
	if loc.row<0 || loc.row>=row{
		return 0,false
	}
	if loc.col<0 || loc.col >= col{
		return 0, false
	}
	return m[loc.row][loc.col],true
}

func walk(maze [][]int, start, end point) [][]int{
	row := len(maze)
	col := len(maze[0])
	// initial path matrix
	path := make([][]int, row)
	for i:=range path{
		path[i]=make([]int,col)
	}
	// queue
	q:= []point{
		{0,0},
	}
	for len(q)>0{
		cur := q[0]
		q=q[1:]
		for _,dir:=range directions{
			next := cur.add(dir)
			// cur is endpoint
			if cur==end{
				return path
			}
			// valid in maze
			valueMaze,ok:=getValue(maze,next)
			if ok==false|| valueMaze==1{
				continue
			}
			// valid in path
			valuePath,ok:=getValue(path, next)
			if ok==false|| valuePath!=0{
				continue
			}
			// not start point
			if next==start{
				continue
			}

			distanceCur,_:=getValue(path,cur)
			path[next.row][next.col]=distanceCur+1
			q=append(q,next)
		}
	}
	return path
}

func getPath(path [][]int, end, start point)[]point{
	ret:= []point{end}
	cur := end
	for cur != start{
		curValue,_:=getValue(path,cur)
		for _,dir:=range directions{
			next := cur.add(dir)
			nextValue, ok:=getValue(path,next)
			if !ok{
				continue
			}
			if nextValue == curValue-1{
				ret = append(ret, next)
				cur = next
				break
			}
		}
	}
	return ret
}

func main() {
	maze := readMaze("maze/maze.txt")
	path:=walk(maze, point{0,0},point{len(maze)-1,len(maze[0])-1})
	printMaze(path)
	arr:=getPath(path,point{len(maze)-1,len(maze[0])-1},point{0,0})
	for _,v:=range arr{
		fmt.Printf("(%d,%d),",v.row,v.col)
	}
	fmt.Println()
}
