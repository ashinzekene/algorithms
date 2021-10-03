package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PassingCars(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{1, 0, 1, 1, 1, 1},
			4,
		},
		{
			[]int{0, 0, 0},
			0,
		},
		{
			[]int{1, 1},
			0,
		},
		{
			[]int{0, 0, 1, 1},
			4,
		},
		{
			[]int{0, 1, 0, 1, 1},
			5,
		},
	}
	for _, tc := range testCases {
		actual := PassingCars(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
