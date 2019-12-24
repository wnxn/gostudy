package hammingdistance

func hammingDistance(x int, y int) int {
	z := x^y
	res :=0
	for z >0{
		z&=z-1
		res++
	}
	return res
}
