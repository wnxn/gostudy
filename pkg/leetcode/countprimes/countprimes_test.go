package countprimes

import (
	"fmt"
	"testing"
)

//func countPrimes(n int) int {
//	arr := make([]bool, n)
//	for i:=range arr{
//		arr[i]=true
//	}
//
//	for i:=2;i<n;i++{
//		if arr[i]{
//			arr[i]=true
//			for j:=i*i;j<n;j+=i{
//				arr[j]=false
//			}
//		}
//	}
//	res :=0
//	for i:=2;i<n;i++{
//		if arr[i]==true{
//			res++
//		}
//	}
//	return res
//}

func countPrimes(n int)int{
	count :=0
	arr := make([]bool,n)
	for i:=2;i<n;i++{
		if !arr[i]{
			count++
			for j:=i*i;j<n;j+=i{
				arr[j]=true
			}
		}
	}
	return count
}

func isPowerOfThree(n int) bool {
	if n == 1{
		return true
	}
	return (n%3==0) &&isPowerOfThree(n/3)
}

func isPrime(n int)bool{
	if n == 1{
		return false
	}
	for i:=2;i*i<=n;i++{
		if n%i == 0{
			return false
		}
	}
	return true
}

func TestCountPrimes(t *testing.T){
	res := countPrimes(1500000)
	fmt.Println(res)
}
