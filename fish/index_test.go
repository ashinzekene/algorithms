package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenomicRangeQuery(t *testing.T) {
	testCases := []struct {
		A        []int
		B        []int
		Expected int
	}{
		{
			[]int{2, 5, 0},
			[]int{1, 1, 1},
			3,
		},
		{
			[]int{4, 3, 2, 1, 5},
			[]int{0, 1, 0, 0, 0},
			2,
		},
		{
			[]int{1, 50, 12, 5, 3, 4, 5, 48},
			[]int{0, 1, 1, 0, 1, 1, 1, 0},
			2,
		},
		{
			[]int{1, 15, 12, 5, 3, 4, 5, 48},
			[]int{1, 1, 1, 0, 1, 1, 1, 0},
			1,
		},
	}
	for _, tc := range testCases {
		result := Fish(tc.A, tc.B)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for A: %v and B: %v", tc.A, tc.B))
	}
}
