package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_distributeCandies(t *testing.T) {
	testCases := []struct {
		Candies  []int
		Expected int
	}{
		{
			[]int{1, 2, 3, 4, 10, 20, 30, 40, 100, 200},
			5,
		},
		{
			[]int{1, 2, 1, 2, 1},
			2,
		},
		{
			[]int{1, 1, 2, 2, 3, 3},
			3,
		},
		{
			[]int{1, 1, 2, 3},
			2,
		},
	}
	for _, tc := range testCases {
		actual := distributeCandies(tc.Candies)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Candies))
	}
}
