package algorithms

func PermutationEquation(p []int32) []int32 {
	result := make([]int32, len(p))
	d := make([]int, len(p)+1)
	// Prepare dictionary
	for index, item := range p {
		d[item] = index
	}

	for i := range p {
		result[i] = int32(d[d[i+1]+1] + 1)
	}
	return result
}
