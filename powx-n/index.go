package algorithms

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 1 {
		return myPow(1/x, n*-1)
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)
}
