package algorithms

func FrogRiverOne(X int, A []int) int {
	leavesPositions := make([]bool, X+1)
	total := 0
	for seconds, leavePosition := range A {
		if !leavesPositions[leavePosition] {
			leavesPositions[leavePosition] = true
			total += 1
		}
		if total == X {
			return seconds
		}
	}
	return -1
}
