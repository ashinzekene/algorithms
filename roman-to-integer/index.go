package algorithms

func romanToInt(s string) int {
	rInts := map[byte]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		current := rInts[s[i]]
		if i < len(s)-1 && rInts[s[i+1]] > current {
			current = rInts[s[i+1]] - current
			i++
		}
		result += current
	}
	return result
}

// func romanToInt(s string) int {
// 	remainderInts := map[string]int{
// 		"C": 100,
// 		"X": 10,
// 		"I": 1,
// 	}
// 	romanInts := map[string]int{
// 		"M": 1000,
// 		"D": 500,
// 		"C": 100,
// 		"L": 50,
// 		"X": 10,
// 		"V": 5,
// 		"I": 1,
// 	}
// 	number := 0
// 	temp := 0
// 	for _, char := range s {
// 		str := string(char)
// 		result := romanInts[str]
// 		if temp != 0 && temp < result {
// 			result -= (temp * 2)
// 		}
// 		temp = remainderInts[str]
// 		number += result
// 	}

// 	return number
// }
