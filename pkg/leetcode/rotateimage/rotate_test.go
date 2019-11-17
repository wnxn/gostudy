package rotateimage

import (
	"fmt"
	"reflect"
	"testing"
)

func rotate(matrix [][]int)  {
	length := len(matrix)
	for i:=0;i< (length)/2;i++{
		for j:=0;j< (length+1)/2;j++{
			tmp := matrix[i][j]
			matrix[i][j]=matrix[length-1-j][i]
			matrix[length-1-j][i]=matrix[length-1-i][length-1-j]
			matrix[length-1-i][length-1-j]=matrix[j][length-1-i]
			matrix[j][length-1-i]=tmp
			fmt.Printf("%d,%d %v\n", i,j,matrix)
		}

	}

}


func TestRotate(t *testing.T){
	tests := []struct{
		matrix [][]int
		res [][]int
	}{
		{
			matrix:[][]int{
				{1,2,3},
				{4,5,6},
				{7,8,9},
			},
			res: [][]int{
				{7,4,1},
				{8,5,2},
				{9,6,3},
			},
		},
		{
			matrix:[][]int{
				{5,1,9,11},
				{2,4,8,10},
				{13,3,6,7},
				{15,14,12,16},
			},
			res: [][]int{
				{15,13,2,5},
				{14,3,4,1},
				{12,6,8,9},
				{16,7,10,11},
			},
		},
	}
	for _, test := range tests{
		rotate(test.matrix)
		if !reflect.DeepEqual(test.matrix, test.res){
			t.Errorf("expect %v, but actually %v", test.res, test.matrix)
		}
	}
}
