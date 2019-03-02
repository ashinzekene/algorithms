package main

import "fmt"

func main() {
	fmt.Println(countingValleys(1, "UDDDUDUU"))
}

func countingValleys(n int32, s string) int32 {
	counter := 0
	var result int32 = 0
	for i := 0; i < len(s); i++ {
		if s[i] == D {
			counter--
		} else {
			counter++
			if counter == 0 {
				result++
			}
		}
	}
	return result
}
