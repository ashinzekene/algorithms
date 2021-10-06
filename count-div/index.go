package algorithms

func CountDiv(A int, B int, K int) int {
	divB := B / K
	divA := A / K
	result := divB - divA
	if A%K == 0 && A != 0 {
		return result + 1
	}
	return result
}
