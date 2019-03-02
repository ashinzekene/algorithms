package main

func main() {
	pageCount(40, 4)
}

func pageCount(n int32, p int32) int32 {
	if n == p || p == 1 {
		return 0
	}
	if n/2 >= p {
		return p / 2
	} else {
		if n%2 == 0 && p%2 == 1 {
			return ((n - p) / 2) + 1
		}
		return (n - p) / 2
	}
}
