package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LargestRectangle(t *testing.T) {
	testCases := []struct {
		Heights  []int32
		Expected int64
	}{
		{
			[]int32{1, 2, 3, 4, 5},
			9,
		},
		{
			[]int32{1, 2, 3, 4, 5, 3, 3, 2},
			15,
		},
		{
			[]int32{11, 11, 10, 10, 10},
			50,
		},
		{
			[]int32{8979, 4570, 6436, 5083, 7780, 3269, 5400, 7579, 2324, 2116},
			26152,
		},
	}
	for _, tc := range testCases {
		actual := largestRectangle(tc.Heights)
		assert.Equal(t, tc.Expected, actual)
	}
}
