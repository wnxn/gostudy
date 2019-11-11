package myatoi

func MyAtoi(str string) int {
	res, len, symbol, breaked := 0, 0, 1, false
	hasValue := false
	for _, v := range str {
		if breaked == true {
			break
		}
		if symbol*res < -1*0x80000000 {
			return -1 * 0x80000000
		}
		if symbol*res > 0x7FFFFFFF {
			return 0x7FFFFFFF
		}
		switch {
		case v >= '0' && v <= '9':
			//fmt.Println(res, v, len,v)
			if res == 0 {
				res = int(v - '0')
			} else {
				res = 10*res + int(v-'0')
			}

			len++
			hasValue = true
		case v == ' ':
			if hasValue == true {
				breaked = true
			}
			continue
		case v == '-':
			if hasValue {
				breaked = true
				break
			}
			symbol = -1
			hasValue = true
		case v == '+':
			if hasValue {
				breaked = true
				break
			}
			symbol = 1
			hasValue = true
		default:
			if hasValue == false {
				breaked = true
			} else {
				breaked = true
			}
		}
	}
	res = res * symbol
	if res < -1*0x80000000 {
		return -1 * 0x80000000
	}
	if res > 0x7FFFFFFF {
		return 0x7FFFFFFF
	}
	return res
}
