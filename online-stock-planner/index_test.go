package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StockSpanner(t *testing.T) {
	testCases := []struct {
		Array    []int
		Expected []int
	}{
		{
			[]int{100, 80, 60, 70, 60, 75, 85},
			[]int{1, 1, 1, 2, 1, 4, 6},
		},
		{
			[]int{60, 70, 80, 90, 100, 50},
			[]int{1, 2, 3, 4, 5, 1},
		},
		{
			[]int{100, 70, 120},
			[]int{1, 1, 3},
		},
		{
			[]int{10},
			[]int{1},
		},
	}
	for _, tc := range testCases {
		ss := Constructor()
		actual := make([]int, len(tc.Array))
		for i, stock := range tc.Array {
			actual[i] = ss.Next(stock)
		}
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Array))
	}
}
