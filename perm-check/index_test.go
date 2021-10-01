package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PermCheck(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{4, 1, 3, 2},
			1,
		},
		{
			[]int{4, 1, 3},
			0,
		},
		{
			[]int{1, 2, 2, 4},
			0,
		},
	}
	for _, tc := range testCases {
		actual := PermCheck(tc.Nums)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
