package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jumpGame(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected bool
	}{
		{
			[]int{1, 2, 3, 5},
			true,
		},
		{
			[]int{4, 5, 2, 1, 0, 4},
			true,
		},
		{
			[]int{4, 0, 2, 1, 0, 4},
			false,
		},
		{
			[]int{4, 2, 2, 0, 0, 4},
			false,
		},
		{
			[]int{3},
			true,
		},
		{
			[]int{1, 2},
			true,
		},
		{
			[]int{0},
			true,
		},
		{
			[]int{4, 0},
			true,
		},
		{
			[]int{0, 4},
			false,
		},
	}
	for _, tc := range testCases {
		actual := canJump(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Cost %d", tc.Nums))
	}
}
