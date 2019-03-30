package algorithms

import (
	"fmt"
	"strings"
)

func main() {
	var s string = "a"
	var n int64 = 1000000000000
	fmt.Println(repeatedString(s, n))
}

func repeatedString(s string, n int64) int64 {
	parts := n / int64(len(s))
	countPerPart := int64(strings.Count(s, "a"))
	restLength := n % int64(len(s))
	remainder := int64(strings.Count(s[: restLength], "a")) 
	return parts*countPerPart + remainder
}