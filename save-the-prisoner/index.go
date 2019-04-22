package algorithms

func SaveThePrisoner(n int32, m int32, s int32) int32 {
	no_of_sweets := m % n
	return s + no_of_sweets - 1
}
