package reverse

func Reverse1(x int) int {
	symbol := 1
	if x < 0 {
		symbol = -1
		x = symbol * x
	}
	arr := []int{}
	for x > 0 {
		arr = append(arr, x%10)
		x = x / 10
	}
	res := 0
	for _, v := range arr {
		res += v
		res *= 10
	}

	res = symbol * res / 10
	if res < -1*0x7FFFFFFF || res > 0x7FFFFFFE {
		res = 0
	}
	return res
}

func Reverse(x int) int {
	symbol := 1
	if x < 0 {
		symbol = -1
		x = symbol * x
	}
	res := 0
	for x > 0 {
		v := x % 10
		x = x / 10
		res += v
		res *= 10
	}

	res = symbol * res / 10
	if res < -1*0x7FFFFFFF || res > 0x7FFFFFFE {
		res = 0
	}
	return res

}
