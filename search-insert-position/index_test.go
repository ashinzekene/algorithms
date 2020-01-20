package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseInteger(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Target   int
		Expected int
	}{
		{
			[]int{1, 3, 5, 6},
			5,
			2,
		},
		{
			[]int{1, 3, 5, 6},
			2,
			1,
		},
		{
			[]int{1, 3, 5, 6},
			7,
			4,
		},
		{
			[]int{1, 3, 5, 6},
			0,
			0,
		},
	}
	for _, tc := range testCases {
		actual := searchInsert(tc.Nums, tc.Target)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Nums))
	}
}
