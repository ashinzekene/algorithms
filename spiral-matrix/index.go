package algorithms

func spiralOrder(matrix [][]int) []int {
	t := 0
	result := make([]int, 0)
	if len(matrix) == 0 {
		return result
	}
	r := len(matrix[0]) - 1
	b := len(matrix) - 1
	l := 0
	for t <= b && l <= r {
		if b >= t {
			result = append(result, matrix[t][l:r+1]...)
			t++
		}
		if r >= l {
			for i := t; i <= b; i++ {
				result = append(result, matrix[i][r])
			}
			r--
		}
		if b >= t {
			for i := r; i >= l; i-- {
				result = append(result, matrix[b][i])
			}
			b--
		}
		if r >= l {
			for i := b; i >= t; i-- {
				result = append(result, matrix[i][l])
			}
			l++
		}
	}
	return result
}
