package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jumpGame3(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Start    int
		Expected bool
	}{
		{
			[]int{4, 2, 3, 0, 3, 1, 2},
			0,
			true,
		},
		{
			[]int{4, 2, 3, 0, 3, 1, 2},
			5,
			true,
		},
		{
			[]int{3, 0, 2, 1, 2},
			2,
			false,
		},
		{
			[]int{2, 3, 1, 1, 4, 2, 1, 1, 1},
			3,
			false,
		},
		{
			[]int{2, 3, 0, 1, 4},
			2,
			true,
		},
	}
	for _, tc := range testCases {
		actual := canReach(tc.Nums, tc.Start)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Cost %d", tc.Nums))
	}
}
