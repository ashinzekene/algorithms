package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GroupAnagrams(t *testing.T) {
	testCases := []struct {
		Strings  []string
		Expected [][]string
	}{
		{
			[]string{"eat", "tea", "tan", "ate", "nat", "bat"},
			[][]string{
				{"bat"},
				{"tan", "nat"},
				{"eat", "tea", "ate"},
			},
		},
		{
			[]string{""},
			[][]string{
				{""},
			},
		},
		{
			[]string{"ab"},
			[][]string{
				{"ab"},
			},
		},
		{
			[]string{"ab", "cs"},
			[][]string{
				{"ab"},
				{"cs"},
			},
		},
	}
	for _, tc := range testCases {
		actual := groupAnagrams(tc.Strings)
		assert.ElementsMatch(t, tc.Expected, actual, fmt.Sprintf("Failed for groupAnagrams using: %v", tc.Strings))
		actual = groupAnagrams1(tc.Strings)
		assert.ElementsMatch(t, tc.Expected, actual, fmt.Sprintf("Failed for groupAnagrams1 using: %v", tc.Strings))
	}
}
