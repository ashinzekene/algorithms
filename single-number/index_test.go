package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValid(t *testing.T) {
	testCases := []struct {
		Nums     []int
		Expected int
	}{
		{
			[]int{2, 2, 1},
			1,
		},
		{
			[]int{4, 1, 2, 1, 2},
			4,
		},
		{
			[]int{2, 4, 4, 6, 6},
			2,
		},
	}
	for _, tc := range testCases {
		actual := singleNumber(tc.Nums)
		assert.Equal(t, tc.Expected, actual)
	}
}
