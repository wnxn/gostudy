package reversebits

func reverseBits(num uint32) uint32 {
	res := uint32(0)
	for  i:= 0; i<32; i++{
		if num&uint32(1<< uint32(i)) !=0{
			// true
			res |= 1<< uint32(32-i)
		}
	}
	return res
}

func generate(numRows int) [][]int {

	res := make([][]int,numRows)
	if numRows == 0{
		return res
	}
	res[0] = []int{1}
	for i:= 2;i<=numRows;i++{
		cur := make([]int, i)
		cur[0]=1
		for j:=1;j<i-1;i++{
			cur[j] = res[i-1][j-1]+res[i-1][j]
		}
		cur[i-1]=1
		res[i]=cur
	}
	return res

}