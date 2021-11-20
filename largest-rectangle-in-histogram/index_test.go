package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LargestRectangleArea(t *testing.T) {
	testCases := []struct {
		Heights  []int
		Expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			9,
		},
		{
			[]int{1, 2, 3, 4, 5, 3, 3, 2},
			15,
		},
		{
			[]int{11, 11, 10, 10, 10},
			50,
		},
		{
			[]int{1, 2, 1},
			3,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{2, 1, 2},
			3,
		},
		{
			[]int{0},
			0,
		},
		{
			[]int{},
			0,
		},
		{
			[]int{8979, 4570, 6436, 5083, 7780, 3269, 5400, 7579, 2324, 2116},
			26152,
		},
	}
	for _, tc := range testCases {
		actual := LargestRectangleArea(tc.Heights)
		assert.Equal(t, tc.Expected, actual)
	}
}
