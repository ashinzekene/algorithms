package algorithms

import "math"

func AppendAndDelete(s string, t string, k int32) string {
	len_s := float64(len(s))
	len_t := float64(len(t))
	length := math.Min(len_s, len_t)
	length_same := 0
	if k >= int32(len_s+len_t) {
		return "Yes"
	}
	for i := 0; i < int(length); i++ {
		if s[i] == t[i] {
			length_same = i
		} else {
			break
		}
	}
	to_remove := len(s) - length_same
	to_add := len(t) - length_same
	whatsLeft := int(k) - (to_remove + to_add)
	if whatsLeft%2 == 0 {
		return "Yes"
	}
	return "No"
}
