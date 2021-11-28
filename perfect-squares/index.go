package algorithms

import (
	coinChange "github.com/ashinzekene/algorithms/coin-change"
)

func NumSquares(n int) int {
	squares := make([]int, 0)
	current := 1
	for current*current <= n {
		squares = append(squares, current*current)
		current++
	}
	return coinChange.CoinChange(squares, n)
}
