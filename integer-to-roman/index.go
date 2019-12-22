package algorithms

import "strings"

func intToRoman(num int) string {
	numberals := []string{"I", "V", "X", "L", "C", "D", "M"}
	tensIndex := 8
	quotient := 1000
	result := ""
	for quotient >= 1 {
		rem := num / quotient
		num -= rem * quotient
		quotient /= 10
		tensIndex -= 2
		switch true {
		case rem < 1:
			continue
		case rem < 4:
			result += strings.Repeat(numberals[tensIndex], rem)
		case rem == 4:
			result += numberals[tensIndex] + numberals[tensIndex+1]
		case rem == 5:
			result += numberals[tensIndex+1]
		case rem <= 8:
			result += numberals[tensIndex+1] + strings.Repeat(numberals[tensIndex], rem-5)
		case rem <= 9:
			result += numberals[tensIndex] + numberals[tensIndex+2]
		}
	}
	return result
}
