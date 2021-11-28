package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CoinChange(t *testing.T) {
	testCases := []struct {
		Coins    []int
		Amount   int
		Expected int
	}{
		{
			[]int{186, 419, 83, 408},
			6249,
			20,
		},
		{
			[]int{2},
			3,
			-1,
		},
		{
			[]int{3, 5, 8},
			10,
			2,
		},
		{
			[]int{1, 2, 5},
			11,
			3,
		},
		{
			[]int{1},
			2,
			2,
		},
		{
			[]int{1},
			0,
			0,
		},
	}
	for _, tc := range testCases {
		result := CoinChange(tc.Coins, tc.Amount)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Coins))
	}
}
