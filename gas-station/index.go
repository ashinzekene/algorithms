package algorithms

func CanCompleteCircuit(gas []int, cost []int) int {
	topCostIndex := -1
	cummulativeCost := 0
	topCost := 0
	for i := len(cost) - 1; i >= 0; i-- {
		netCost := gas[i] - cost[i]
		cummulativeCost += netCost
		if cummulativeCost >= topCost {
			topCostIndex = i
			topCost = cummulativeCost
		}
	}
	if cummulativeCost >= 0 {
		return topCostIndex
	}
	return -1
}
