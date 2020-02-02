package algorithms

func uniquePaths(m int, n int) int {
	dpMatrix := make([][]int, m)
	for i := 0; i < len(dpMatrix); i++ {
		dpMatrix[i] = make([]int, n)
		dpMatrix[i][0] = 1
	}
	for i := 0; i < len(dpMatrix[0]); i++ {
		dpMatrix[0][i] = 1
	}
	for i := 1; i < len(dpMatrix); i++ {
		for j := 1; j < len(dpMatrix[0]); j++ {
			dpMatrix[i][j] = dpMatrix[i-1][j] + dpMatrix[i][j-1]
		}
	}
	return dpMatrix[m-1][n-1]
}

//[0,1,1,1]
//[1,2,3,4]
//[1,3,6,10]
//[1,4,10,20]
//[1,5,15,35]
//[1,6,21,56]
//[1,7,28,84]

// 1,1 -> 1
// 2,2 -> 2
// 2,3 -> 3
// 2,4 -> 4
// 3,3 -> 6
// 2,5 -> 5
// 3,4 -> 10
// 3,6 -> 15
// 4,4-> 20
// 3,7 -> 21
