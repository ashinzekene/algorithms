package algorithms

import "math"

func AppendAndDelete(s string, t string, k int32) string {
	len_s := float64(len(s))
	len_t := float64(len(t))
	k_float := float64(k)
	length_diff := math.Abs(len_s - len_t)
	length := math.Min(len_s, len_t)
	length_same := 0

	for i := 0; i < int(length); i++ {
		if s[i] == t[i] {
			length_same = i
		} else {
			break
		}
	}
	expected := (length-float64(length_same)-1)*2 + length_diff
	if expected <= k_float && int(k_float-expected)%2 == 0 || length_same*2+int(expected) <= int(k) {
		return "Yes"
	}
	return "No"
}
