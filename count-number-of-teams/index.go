package algorithms

func NumTeams(rating []int) int {
	result := 0
	for i := 1; i < len(rating)-1; i++ {
		greaterLeft, lessLeft := countGreaterLess(rating[:i], rating[i])
		greaterRight, lessRight := countGreaterLess(rating[i+1:], rating[i])
		desc := greaterLeft * lessRight
		asc := lessLeft * greaterRight
		result += asc + desc
	}
	return result
}

func countGreaterLess(rating []int, v int) (int, int) {
	greater := 0
	less := 0
	for _, val := range rating {
		if val > v {
			greater++
		} else {
			less++
		}
	}
	return greater, less
}
