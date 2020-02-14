package algorithms

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	result := make([]int, l1+l2)
	for i := l1 - 1; i >= 0; i-- {
		sum := make([]int, len(num2)+l1-i)
		val1, _ := strconv.Atoi(string(num1[i]))
		for j := l2 - 1; j >= 0; j-- {
			val2, _ := strconv.Atoi(string(num2[j]))
			multiple := val1 * val2
			add(sum, multiple, l1-i-1+l2-j-1)
		}
		add2(result, sum)
	}
	resultStr := ""
	i := len(result) - 1
	for i >= 0 && result[i] == 0 {
		i--
	}
	for ; i >= 0; i-- {
		if i == len(result)-1 && result[i] == 0 {
			continue
		}
		resultStr += strconv.Itoa(result[i])
	}
	if resultStr == "" {
		return "0"
	}
	return resultStr
}

func add(slice []int, val, index int) {
	slice[index] = slice[index] + val%10
	slice[index+1] = val / 10
}

func add2(result, sum []int) {
	remainder := 0
	for i := 0; i < len(sum); i++ {
		addition := result[i] + sum[i] + remainder
		result[i] = addition % 10
		remainder = addition / 10
	}
}
