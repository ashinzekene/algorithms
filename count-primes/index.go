package algorithms

// https://leetcode.com/problems/count-primes/submissions/

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	nums := make([]bool, n)
	nums[0] = true
	nums[1] = true
	nums[2] = false
	for i := 2; i*i < n; i++ {
		if nums[i] {
			// Not prime
			continue
		}
		for j := i + i; j < n; j += i {
			nums[j] = true
		}
	}
	count := 0
	for _, v := range nums {
		if !v {
			count++
		}
	}
	return count
}
