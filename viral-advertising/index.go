package algorithms

func ViralAdvertising(n int32) int32 {
	shared := 5
	liked := 0

	for n > 0 {
		likes := int(shared / 2)
		shared = likes * 3
		liked += likes
		n--
	}
	return int32(liked)
}
