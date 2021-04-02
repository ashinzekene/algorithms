package algorithms

import . "github.com/ashinzekene/algorithms/utils"

// func minCostClimbingStairs(costs []int) int {
// 	minStairsCost := make([]int, len(costs))
// 	minStairsCost[0] = costs[0]
// 	minStairsCost[1] = costs[1]

// 	for i := 2; i < len(costs); i++ {
// 		cost := costs[i]
// 		minStairsCost[i] = Min(minStairsCost[i-1]+cost, minStairsCost[i-2]+cost)
// 	}
// 	return Min(
// 		minStairsCost[len(costs)-1],
// 		minStairsCost[len(costs)-2],
// 	)
// }

func minCostClimbingStairs(costs []int) int {
	a := costs[0]
	b := costs[1]

	for i := 2; i < len(costs); i++ {
		cost := costs[i]
		minCost := Min(a+cost, b+cost)
		a = b
		b = minCost
	}
	return Min(a, b)
}
