package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseInteger(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			0,
			4,
		},
		{
			[]int{4, 5, 6, 7, 0, 1, 2},
			8,
			-1,
		},
		{
			[]int{},
			8,
			-1,
		},
		{
			[]int{4, 5, 6, 7, 8, 9, 10},
			6,
			2,
		},
		{
			[]int{4, 5},
			6,
			-1,
		},
		{
			[]int{5, 6, 7, 8, 9, 10, 1, 2, 3},
			3,
			8,
		},
		{
			[]int{5, 6, 7, 8, 9, 10, 1, 2, 3},
			5,
			0,
		},
	}
	for _, tc := range testCases {
		actual := search(tc.Nums, tc.Target)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
