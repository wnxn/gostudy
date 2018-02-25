package number

import "fmt"

func BinarySearch(array []int, value int)int{
	if len(array) == 0{
		return -1
	}
	length := len(array)
	mid := length/2
	if array[mid] == value{
		return mid
	}else if array[mid] < value{
		cur :=  BinarySearch(array[mid+1:], value)
		if cur == -1{
			return -1
		}else{
			return cur + mid + 1
		}
	}else{
		return BinarySearch(array[:mid],value)
	}
}

func BinarySearch2(array []int, value int)int{

	a:=0
	b:=len(array)
	var mid int
	for a<b{
		mid = a + (b-a)/2
		if array[mid] == value{
			return mid
		}else if  array[mid] < value{
			a= mid+1
		} else{
			b=mid
		}
	}
	return -1

}

func BinarySearchMain(){
	arr:=[]int{1,2,3,4,5,6}
	fmt.Println(BinarySearch2(arr, 5));
}
