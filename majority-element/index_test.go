package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{2, 3, 2},
			2,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
		{
			[]int{1},
			1,
		},
	}
	for _, tc := range testCases {
		actual := majorityElement(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
