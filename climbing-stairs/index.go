package algorithms

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return n
	}
	a := 1
	b := 2
	for i := 3; i <= n; i++ {
		tempa := a
		a = b
		b = tempa + b
	}
	return b
}
