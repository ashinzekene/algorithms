package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProduct(t *testing.T) {
	testCases := []struct {
		Prices   []int
		Expected int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 5, 4, 3, 2, 1},
			0,
		},
		{
			[]int{4, 5, 6, 1},
			2,
		},
		{
			[]int{2, 3, 7},
			5,
		},
		{
			[]int{2, 1, 7},
			6,
		},
		{
			[]int{4, 7},
			3,
		},
	}
	for _, tc := range testCases {
		result := maxProfit(tc.Prices)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Prices))
	}
}
