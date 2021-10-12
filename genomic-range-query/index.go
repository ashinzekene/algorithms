package algorithms

func GenomicRangeQuery(S string, P []int, Q []int) []int {
	geneMap := map[rune]int{
		'A': 1,
		'C': 2,
		'G': 3,
		'T': 4,
	}
	occurenceMap := make([][5]int, len(S)+1)
	result := make([]int, len(P))
	for i, v := range S {
		geneInd := geneMap[v]
		occurenceMap[i+1][1] = occurenceMap[i][1]
		occurenceMap[i+1][2] = occurenceMap[i][2]
		occurenceMap[i+1][3] = occurenceMap[i][3]
		occurenceMap[i+1][4] = occurenceMap[i][4]
		occurenceMap[i+1][geneInd] += 1
	}
	for i := range P {
		start, end := P[i], Q[i]
		for geneIndex := range occurenceMap[0] {
			if occurenceMap[end+1][geneIndex]-occurenceMap[start][geneIndex] > 0 {
				result[i] = geneIndex
				break
			}
		}
	}
	return result
}
