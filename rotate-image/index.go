package algorithms

func rotate(matrix [][]int) {
	length := len(matrix)
	iterations := length / 2
	l := length
	for i := 0; i < iterations; i++ {
		for j := 0; j < l-1; j++ {
			start := i
			end := length - i - 1
			a := matrix[start][start+j]
			b := matrix[start+j][end]
			c := matrix[end][end-j]
			d := matrix[end-j][start]
			matrix[start][start+j] = d
			matrix[start+j][end] = a
			matrix[end][end-j] = b
			matrix[end-j][start] = c
		}
		l -= 2
	}
}
