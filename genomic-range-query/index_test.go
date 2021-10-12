package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GenomicRangeQuery(t *testing.T) {
	testCases := []struct {
		A        string
		B        []int
		K        []int
		Expected []int
	}{
		{
			"CAGCCTA",
			[]int{2, 5, 0},
			[]int{4, 5, 6},
			[]int{2, 4, 1},
		},
		{
			"TAAGGC",
			[]int{1, 3, 0},
			[]int{4, 3, 4},
			[]int{1, 3, 1},
		},
		{
			"ACCCTCCG",
			[]int{2, 5, 0},
			[]int{4, 5, 6},
			[]int{2, 2, 1},
		},
	}
	for _, tc := range testCases {
		result := GenomicRangeQuery(tc.A, tc.B, tc.K)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for query: %s", tc.A))
	}
}
