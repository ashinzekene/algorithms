package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FreqQuery(t *testing.T) {
	testCases := []struct {
		Queries  [][]int32
		Expected []int32
	}{
		{
			[][]int32{
				[]int32{1, 3},
				[]int32{2, 3},
				[]int32{3, 2},
				[]int32{1, 4},
				[]int32{1, 5},
				[]int32{1, 5},
				[]int32{1, 4},
				[]int32{3, 2},
				[]int32{2, 4},
				[]int32{3, 2},
			},
			[]int32{0, 1, 1},
		},
		{
			[][]int32{
				[]int32{1, 5},
				[]int32{1, 6},
				[]int32{3, 2},
				[]int32{1, 10},
				[]int32{1, 10},
				[]int32{1, 6},
				[]int32{2, 5},
				[]int32{3, 2},
			},
			[]int32{0, 1},
		},
	}
	for _, tc := range testCases {
		actual := FreqQuery(tc.Queries)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Digit %d", tc.Queries))
	}
}
