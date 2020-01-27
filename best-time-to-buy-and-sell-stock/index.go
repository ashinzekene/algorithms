package algorithms

import . "github.com/ashinzekene/algorithms/utils"

func maxProfit(prices []int) int {
	maxDiff := 0
	if len(prices) == 0 {
		return 0
	}
	minVal := prices[0]
	for _, v := range prices {
		minVal = Min(v, minVal)
		diff := v - minVal
		maxDiff = Max(maxDiff, diff)
	}
	return maxDiff
}
