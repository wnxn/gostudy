package fizzbuzz

import (
	"reflect"
	"strconv"
	"testing"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i:=1;i <=n;i++{
		if i%3==0 && i%5==0{
			res[i-1]= "FizzBuzz"
		}else if i%3==0{
			res[i-1]="Fizz"
		}else if i%5==0{
			res[i-1]="Buzz"
		}else{
			res[i-1]=strconv.Itoa(i)
		}
	}

	return res
}

func TestFizz(t *testing.T){
	n := 15
	res := []string{
		"1","2", "Fizz","4", "Buzz", "Fizz","7","8",
		"Fizz", "Buzz", "11","Fizz","13","14","FizzBuzz",
	}
	out := fizzBuzz(n)
	if !reflect.DeepEqual(out, res){
		t.Errorf("expect %v, but actually %v", res, out)
	}
}

