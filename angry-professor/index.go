package main

import "fmt"

func main() {
	fmt.Printf("Expected %s, got %s\n", "NO", angryProfessor(4, []int32{-93, -86, 49, -62, -90, -63, 40, 72, 11, 67}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(10, []int32{23, -35, -2, 58, -67, -56, -42, -73, -19, 37}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(9, []int32{13, 91, 56, -62, 96, -5, -84, -36, -46, -13}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(8, []int32{45, 67, 64, -50, -8, 78, 84, -51, 99, 60}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(7, []int32{26, 94, -95, 34, 67, -97, 17, 52, 1, 86}))
	fmt.Printf("Expected %s, got %s\n", "NO", angryProfessor(2, []int32{19, 29, -17, -71, -75, -27, -56, -53, 65, 83}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(10, []int32{-32, -3, -70, 8, -40, -96, -18, -46, -21, -79}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(9, []int32{-50, 0, 64, 14, -56, -91, -65, -36, 51, -28}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(3, []int32{-1, -3, 4, 2}))
	fmt.Printf("Expected %s, got %s\n", "NO", angryProfessor(2, []int32{-1, -3, 4, 2}))
	fmt.Printf("Expected %s, got %s\n", "YES", angryProfessor(5, []int32{-1, -3, 4, 2, -3, 1}))
}

func angryProfessor(k int32, a []int32) string {
	var count int32 = 0
	for _, x := range a {
		if x <= 0 {
			count++
		}
		if count >= k {
			return "NO"
		}
	}
	return "YES"
}
