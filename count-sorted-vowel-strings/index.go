package algorithms

// Intuition
// -	 a 	e	 i  o  u
// n=1  1  1  1  1  1
// n=2  5  4  3  2  1
// n=3  15 10 6  3  1
// n at any vowel = n-1 at that vowel + n at previous vowel
// Eg where n is 3 and vowel is u = n=2 which is 3 + n=3 at i = 3 = 6

func countVowelStrings(n int) int {
	// No of occurrence if n = 1
	vowelCount := []int{1, 1, 1, 1, 1}
	for c := 1; c <= n; c++ {
		for i := 1; i < len(vowelCount); i++ {
			vowelCount[i] = vowelCount[i] + vowelCount[i-1]
		}
	}
	return vowelCount[len(vowelCount)-1]
}
