package algorithms

import "math"

func TheLoveLetterMystery(s string) int32 {
	var result int32
  length := len(s)
	l := int(math.Ceil(float64(length - 1) / 2))
	for i:=0; i<l; i++ {
		if s[i] == s[length-1-i] {
			continue
		}
		diff := int(s[i]) - int(s[length-i-1])
		result += int32(math.Abs(float64(diff)))
	}
	return result
}