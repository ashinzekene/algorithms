package algorithms

func IsInterleave(s1 string, s2 string, s3 string) bool {
	dpMatrix := make([][]bool, len(s1)+1)
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	for i := range dpMatrix {
		dpMatrix[i] = make([]bool, len(s2)+1)
	}
	dpMatrix[0][0] = true
	for i := 1; i < len(s1)+1; i++ {
		dpMatrix[i][0] = dpMatrix[i-1][0] && s3[i-1] == s1[i-1]
	}
	for i := 1; i < len(s2)+1; i++ {
		dpMatrix[0][i] = dpMatrix[0][i-1] && s3[i-1] == s2[i-1]
	}
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			index := i + j - 1
			if index < len(s3) {
				equalS1 := s3[index] == s1[i-1]
				equalS2 := s3[index] == s2[j-1]
				dpMatrix[i][j] = equalS1 && dpMatrix[i-1][j] || equalS2 && dpMatrix[i][j-1]
			}
		}
	}
	return dpMatrix[len(s1)][len(s2)]
}
