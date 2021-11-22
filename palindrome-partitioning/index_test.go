package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Partition(t *testing.T) {
	testCases := []struct {
		S        string
		Expected [][]string
	}{
		{
			"aab",
			[][]string{
				{"a", "a", "b"},
				{"aa", "b"},
			},
		},
		{
			"a",
			[][]string{
				{"a"},
			},
		},
		{
			"abbab",
			[][]string{
				{"a", "b", "b", "a", "b"},
				{"a", "b", "bab"},
				{"a", "bb", "a", "b"},
				{"abba", "b"},
			},
		},
		{
			"aabcdef",
			[][]string{
				{"a", "a", "b", "c", "d", "e", "f"},
				{"aa", "b", "c", "d", "e", "f"},
			},
		},
		{
			"aabb",
			[][]string{
				{"a", "a", "b", "b"},
				{"aa", "b", "b"},
				{"a", "a", "bb"},
				{"aa", "bb"},
			},
		},
	}
	for _, tc := range testCases {
		actual := Partition(tc.S)
		assert.ElementsMatch(t, tc.Expected, actual, "Partition did not work for "+tc.S)
		actual = Partition1(tc.S)
		assert.ElementsMatch(t, tc.Expected, actual, "Partition1 did not work for "+tc.S)
	}
}
