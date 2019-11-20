package isvalidsudoku

import (
	"testing"
)

//func isValidSudoku(board [][]byte) bool {
//	for i:=0;i<9;i++{
//		var tmp []byte
//		for j:=0;j<9;j++{
//			tmp = append(tmp, board[j][i])
//		}
//		if !isValidRow(tmp){
//			return false
//		}
//	}
//
//
//	for _,v:=range board{
//		if !isValidRow(v){
//			return false
//		}
//	}
//
//	for i:=0;i<3;i++{
//
//		for j:=0;j<3;j++{
//			var tmp []byte
//			tmp = append(tmp, board[i*3][j*3:j*3+3]...)
//			tmp = append(tmp, board[i*3+1][j*3:j*3+3]...)
//			tmp = append(tmp, board[i*3+2][j*3:j*3+3]...)
//			if !isValidRow(tmp){
//				return false
//			}
//		}
//
//	}
//	return true
//}
//
//func isValidRow(row []byte)bool{
//	fmt.Printf("%c\n",row)
//	mp := make(map[byte]struct{})
//	for _,v:=range row{
//		if v == '.'{
//			continue
//		}
//		if _, ok:=mp[v];ok{
//			return false
//		}else{
//			mp[v]= struct{}{}
//		}
//	}
//	return true
//}

//func isValidSudoku(board [][]byte) bool {
//
//}

func TestIsValidSudoku(t *testing.T) {
	tests := []struct {
		board [][]byte
		res   bool
	}{
		{
			board: [][]byte{
				{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
				{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
				{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
				{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
				{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
				{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
				{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
				{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
				{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
			},
			res: true,
		},
		//{
		//	board: [][]byte{
		//		{'.','.','4','.','.','.','6','3','.'},
		//		{'.','.','.','.','.','.','.','.','.'},
		//		{'5','.','.','.','.','.','.','9','.'},
		//		{'.','.','.','5','6','.','.','.','.'},
		//		{'4','.','3','.','.','.','.','.','1'},
		//		{'.','.','.','7','.','.','.','.','.'},
		//		{'.','.','.','.','.','.','.','.','.'},
		//		{'.','.','.','.','.','.','.','.','.'},
		//		{'.','.','4','.','.','.','.','6','3'},
		//	},
		//	res: false,
		//},
	}
	for _, test := range tests {
		res := isValidSudoku(test.board)
		if res != test.res {
			t.Errorf("expect %t, but actually %t", test.res, res)
		}
	}
}
