package algorithms

func arrayManipulation(n int32, queries [][]int32) {
	maxIndexes := make([]int32, 1, n)
	arr := make([]int32, n, n)
	for _, q := range queries {
		for i := q[0] - 1; i < q[1]; i++ {
			v := arr[i] + q[2]
			if v == arr[maxIndexes[0]] {
				maxIndexes = append(maxIndexes, i)
			}
		}
	}

}
