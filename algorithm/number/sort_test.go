package number

import "testing"

var testcases = []struct{
	array []int
	result []int
}{
	{[]int{23,34,45,12,54,78,45,98}, []int{12,23,34,45,45,54,78,98}},
}

func TestMergeSort(t *testing.T) {
	for _,v:=range testcases{
		res :=MergeSort(v.array)
		for i,_:=range res{
			if res[i]!= v.result[i]{
				t.Errorf("TestMergeSort eroor expect %v, but result is %v", v.result, res)
			}
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i:=0;i<b.N;i++{
		MergeSort(testcases[0].array)
	}
}