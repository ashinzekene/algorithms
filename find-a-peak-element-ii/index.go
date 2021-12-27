package algorithms

func findPeakGrid(mat [][]int) []int {
	l := 0
	r := len(mat) - 1
	i := 0
	for l < r {
		m := l + (r-l)/2
		i = getMaxIndex(mat[m])
		if mat[m][i] < mat[m+1][i] {
			l = m + 1
		} else {
			r = m
		}
	}
	return []int{l, i}
}

func getMaxIndex(row []int) int {
	max := row[0]
	maxIndex := 0
	for i := 0; i < len(row); i++ {
		if row[i] > max {
			max = row[i]
			maxIndex = i
		}
	}
	return maxIndex
}
