package algorithms

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		diff := prices[i] - prices[i-1]
		if diff > 0 {
			profit += diff
		}
	}
	return profit
}
