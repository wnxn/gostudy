package number

func MergeSort(arr []int)[]int {
	length := len(arr)
	if length <= 1 {
		return arr
	}
	mid := length / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	llen := len(left)
	rlen := len(right)
	var ret []int
	lindex :=0
	rindex:=0
	for lindex < llen && rindex < rlen{
		if left[lindex] < right[rindex]{
			ret = append(ret, left[lindex])
			lindex++
		}else{
			ret = append(ret, right[rindex])
			rindex++
		}
	}
	ret = append(ret, left[lindex:]...)
	ret = append(ret, right[rindex:]...)
	return ret
}