package algorithms

import (
	"sort"

	. "github.com/ashinzekene/algorithms/utils"
)

func CoinChange(coins []int, amount int) int {
	sort.Ints(coins)
	dp := make([]int, amount+1)

	for currentAmount := 1; currentAmount <= amount; currentAmount++ {
		dp[currentAmount] = -1
		for _, coin := range coins {
			if coin > currentAmount {
				break
			}
			if dp[currentAmount-coin] >= 0 {
				if dp[currentAmount] == -1 {
					dp[currentAmount] = dp[currentAmount-coin] + 1
				} else {
					dp[currentAmount] = Min(dp[currentAmount], dp[currentAmount-coin]+1)
				}
			}
		}
	}
	return dp[amount]
}
