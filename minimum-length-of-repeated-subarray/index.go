package algorithms

func findLength(nums1 []int, nums2 []int) int {
	l1 := len(nums1)
	l2 := len(nums2)
	max := 0
	prevDp := make([]int, l2)
	for i := 0; i < l1; i++ {
		dp := make([]int, l2)
		for j := 0; j < l2; j++ {
			num1 := nums1[i]
			num2 := nums2[j]
			result := 0
			if num1 == num2 {
				result = 1
			}
			if i > 0 && j > 0 && result > 0 {
				result += prevDp[j-1]
			}
			dp[j] = result
			if dp[j] > max {
				max = dp[j]
			}
		}
		prevDp = dp
	}
	return max
}
