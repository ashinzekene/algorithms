package algorithms

import (
	"strconv"
)

func RestoreIpAddresses(s string) []string {
	result := make([]string, 0)
	computeAddress(&result, "", 4, s)
	for i := range result {
		result[i] = result[i][1:]
	}

	return result
}

func computeAddress(result *[]string, ipaddr string, intsLeft int, strLeft string) {
	if intsLeft == 0 && strLeft == "" {
		*result = append(*result, ipaddr)
		return
	}
	if !canBeValidAdressRange(strLeft, intsLeft) {
		return
	}
	int1, _ := strconv.Atoi(strLeft[0:1])
	computeAddress(result, ipaddr+"."+strconv.Itoa(int1), intsLeft-1, strLeft[1:])
	if strLeft[0] == '0' {
		return
	}
	if len(strLeft) > 1 {
		int2, _ := strconv.Atoi(strLeft[0:2])
		computeAddress(result, ipaddr+"."+strconv.Itoa(int2), intsLeft-1, strLeft[2:])
	}
	if len(strLeft) > 2 {
		int3, _ := strconv.Atoi(strLeft[0:3])
		if int3 <= 255 {
			computeAddress(result, ipaddr+"."+strconv.Itoa(int3), intsLeft-1, strLeft[3:])
		}
	}
}

func canBeValidAdressRange(str string, intLength int) bool {
	return intLength <= len(str) && len(str) <= (3*intLength)
}
