package algorithms

func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	var pos int32 = 1
	currentIndex := 0
	result := make([]int32, len(alice))
	for i := len(alice) - 1; i >= 0; i-- {
		aliceScore := alice[i]
		for j := currentIndex; j < len(scores); j++ {
			score := scores[j]
			if aliceScore >= score {
				result[i] = pos
				currentIndex = j
				break
			} else if j == len(scores)-1 {
				for k := i; k <= 0; k++ {
					result[k] = pos + 1
				}
				return result
			} else if score > scores[j+1] {
				pos++
			}
			currentIndex = j
		}
	}
	return result
}
