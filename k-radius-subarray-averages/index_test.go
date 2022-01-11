package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_jumpGame3(t *testing.T) {
	testCases := []struct {
		K        int
		Nums     []int
		Expected []int
	}{
		{
			3,
			[]int{7, 4, 3, 9, 1, 8, 5, 2, 6},
			[]int{-1, -1, -1, 5, 4, 4, -1, -1, -1},
		},
		{
			0,
			[]int{10000},
			[]int{10000},
		},
		{
			100,
			[]int{8},
			[]int{-1},
		},
		{
			3,
			[]int{12, 28, 83, 4, 25, 26, 25, 2, 25, 25, 25, 12},
			[]int{-1, -1, -1, 29, 27, 27, 18, 21, 20, -1, -1, -1},
		},
	}
	for _, tc := range testCases {
		actual := getAverages(tc.Nums, tc.K)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Cost %d", tc.Nums))
	}
}
