package algorithms

import "strconv"

func monotoneIncreasingDigits(n int) int {
	nStr := strconv.Itoa(n)
	nList := make([]int, len(nStr))
	for i, v := range nStr {
		nList[i], _ = strconv.Atoi(string(v))
	}
	result := ""
	for i := len(nList) - 1; i > 0; i-- {
		currentInt := nList[i]
		nextInt := nList[i-1]
		if currentInt < nextInt {
			newResult := ""
			for range result {
				newResult += "9"
			}
			result = newResult
			result = "9" + result
			nList[i-1] -= 1
		} else {
			result = strconv.Itoa(currentInt) + result
		}
	}
	first := nList[0]
	result = strconv.Itoa(first) + result
	resultInt, _ := strconv.Atoi(result)
	return resultInt
}
