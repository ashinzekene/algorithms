package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wiggleMaxLength(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{1, 7, 4, 9, 2, 5},
			6,
		},
		{
			[]int{1, 17, 5, 15, 10, 16, 8},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			2,
		},
		{
			[]int{0, 0},
			1,
		},
		{
			[]int{1, 0, 0},
			2,
		},
		{
			[]int{1, 6, 6},
			2,
		},
		{
			[]int{1, 1, 1, 1, 1, 1},
			1,
		},
	}
	for _, tc := range testCases {
		actual := wiggleMaxLength(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
