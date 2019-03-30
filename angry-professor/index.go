package algorithms

func AngryProfessor(k int32, a []int32) string {
	var count int32 = 0
	for _, x := range a {
		if x <= 0 {
			count++
		}
		if count >= k {
			return "NO"
		}
	}
	return "YES"
}
