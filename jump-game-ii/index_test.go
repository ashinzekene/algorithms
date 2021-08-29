package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jumpGame2(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{2, 3, 1, 1, 4},
			2,
		},
		{
			[]int{2, 2, 4, 1, 4},
			2,
		},
		{
			[]int{2, 2, 1, 1, 4},
			3,
		},
		{
			[]int{2, 3, 1, 1, 4, 2, 1, 1, 1},
			3,
		},
		{
			[]int{2, 3, 0, 1, 4},
			2,
		},
	}
	for _, tc := range testCases {
		actual := jump(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Cost %d", tc.Nums))
	}
}
