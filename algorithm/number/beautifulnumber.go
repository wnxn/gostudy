package number

import "fmt"

func getRule(value int)int{
	var rule int
	for rule = 2; rule < value; rule++{
		cur := value
		for {
			if cur == 1{
				return rule
			}
			if cur%rule == 1{
				cur /=rule
			}else{
				break
			}
		}
	}
	return rule
}

func BeautifulNumber(){
	for i:=3;i<20;i++{
		fmt.Printf("Number %d: (%d)\n", i, getRule(i))
	}
}
