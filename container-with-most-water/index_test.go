package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxArea(t *testing.T) {
	testCases := []struct {
		Height   []int
		Expected int
	}{
		{
			[]int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			49,
		},
		{
			[]int{1, 2, 3, 4, 5, 3, 3, 2},
			12,
		},
		{
			[]int{11, 11, 10, 10, 10},
			40,
		},
		{
			[]int{2, 3, 4, 5, 18, 17, 6},
			17,
		},
		{
			[]int{1, 2, 1},
			2,
		},
	}
	for _, tc := range testCases {
		result := maxArea(tc.Height)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Height))
	}
}
