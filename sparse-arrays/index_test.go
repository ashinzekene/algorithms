package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SparseArrays(t *testing.T) {
	testCases := []struct {
		Strings  []string
		Queries  []string
		Expected []int32
	}{
		{
			[]string{"aba", "baba", "aba", "xzxb"},
			[]string{"aba", "xzxb", "ab"},
			[]int32{2, 1, 0},
		},
		{
			[]string{"def", "de", "fgh"},
			[]string{"de", "lmn", "fgh"},
			[]int32{1, 0, 1},
		},
		{
			[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn", "sdaklfj", "asdjf", "na", "asdjf", "na", "basdn", "sdaklfj", "asdjf"},
			[]string{"abcde", "sdaklfj", "asdjf", "na", "basdn"},
			[]int32{1, 3, 4, 3, 2},
		},
	}
	for _, tc := range testCases {
		actual := matchingStrings(tc.Strings, tc.Queries)
		assert.Equal(t, tc.Expected, actual)
	}
}
