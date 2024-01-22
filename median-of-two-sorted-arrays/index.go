package algorithms

import (
	"github.com/ashinzekene/algorithms/utils"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	L := len1 + len2
	m := L / 2
	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}
	r := len1 - 1
	l := 0
	for {
		m1 := (l + r) / 2
		m2 := m - m1 - 2

		l1 := -10 ^ 7
		r1 := 10 ^ 7
		l2 := -10 ^ 7
		r2 := 10 ^ 7

		if m1 >= 0 && m1 < len(nums1) {
			l1 = nums1[m1]
		}
		if m1+1 < len(nums1) {
			r1 = nums1[m1+1]
		}
		if m2 >= 0 && m2 < len(nums2) {
			l2 = nums2[m2]
		}
		if m2+1 < len(nums2) {
			r2 = nums2[m2+1]
		}

		if l1 <= r2 && l2 <= r1 {
			if L%2 == 1 {
				return float64(utils.Min(r1, r2))
			}
			return float64(utils.Max(l1, l2) + utils.Min(r1, r2)) / 2
		} else if l1 > r2 {
			r = m1 - 1
		} else {
			l = m1 + 1
		}
	}
}
