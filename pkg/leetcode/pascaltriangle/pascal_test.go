package pascaltriangle

import (
	"reflect"
	"testing"
)

func generate(numRows int) [][]int {

	res := make([][]int,numRows)
	if numRows == 0{
		return res
	}
	res[0] = []int{1}
	for i:= 2;i<= numRows;i++{
		cur := make([]int, i)
		cur[0]=1
		for j:=2;j<i;j++{
			cur[j-1] = res[i-2][j-2]+res[i-2][j-1]
		}
		cur[i-1]=1
		res[i-1]=cur
	}
	return res

}

func TestGenerate(t *testing.T){
	tests :=[]struct{
		rows int
		res [][]int
	}{
		{
			rows:5,
			res: [][]int{
				{1},
				{1,1},
				{1,2,1},
				{1,3,3,1},
				{1,4,6,4,1},
			},
		},
	}
	for _,test:=range tests{
		res := generate(test.rows)
		if !reflect.DeepEqual(res, test.res){
			t.Errorf("expect %v, but actually %v", test.res, res)
		}
	}
}
