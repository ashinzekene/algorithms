package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TrapTainWater(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			6,
		},
	}
	for _, tc := range testCases {
		actual := trap(tc.Nums)
		actual2 := trap2(tc.Nums)
		assert.Equal(t, tc.Expected, actual)
		assert.Equal(t, tc.Expected, actual2)
	}
}
