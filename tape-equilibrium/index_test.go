package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TapeEquilibrium(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{-1, 0, 1, 2, -1, -4},
			1,
		},
		{
			[]int{0, 0, 0},
			0,
		},
		{
			[]int{200, -200},
			400,
		},
		{
			[]int{-3, -1, -2, -4, -3},
			1,
		},
	}
	for _, tc := range testCases {
		actual := TapeEquilibrium(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
