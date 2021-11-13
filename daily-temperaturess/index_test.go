package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DailyTemperatures(t *testing.T) {
	testCases := []struct {
		Temperatures []int
		Expected     []int
	}{
		{
			[]int{73, 74, 75, 71, 69, 72, 76, 73},
			[]int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			[]int{30, 40, 50, 60},
			[]int{1, 1, 1, 0},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{70},
			[]int{0},
		},
		{
			[]int{70, 60},
			[]int{0, 0},
		},
		{
			[]int{70, 60, 100},
			[]int{2, 1, 0},
		},
		{
			[]int{70, 60, 50, 40, 30, 20, 10, 100},
			[]int{7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tc := range testCases {
		result := DailyTemperatures(tc.Temperatures)
		assert.Equal(t, tc.Expected, result, "Failed for %v", tc.Temperatures)
	}
}
