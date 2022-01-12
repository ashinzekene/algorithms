package algorithms

func numOfSubarrays(arr []int, k int, threshold int) int {
	if k > len(arr) {
		return 0
	}
	// first average
	lastSum := 0
	for i := 0; i < k; i++ {
		lastSum += arr[i]
	}
	prevAverage := float64(lastSum) / float64(k)
	result := 0
	if prevAverage >= float64(threshold) {
		result++
	}
	for i := 1; i <= len(arr)-k; i++ {
		sum := lastSum - arr[i-1] + arr[i+k-1]
		average := float64(sum) / float64(k)
		if average >= float64(threshold) {
			result++
		}
		lastSum = sum
	}
	return result
}
