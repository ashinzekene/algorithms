package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountBits(t *testing.T) {
	testCases := []struct {
		N        int
		Expected []int
	}{
		{
			2,
			[]int{0, 1, 1},
		},
		{
			6,
			[]int{0, 1, 1, 2, 1, 2, 2},
		},
		{
			7,
			[]int{0, 1, 1, 2, 1, 2, 2, 3},
		},
		{
			1,
			[]int{0, 1},
		},
		{
			0,
			[]int{0},
		},
	}
	for _, tc := range testCases {
		result := CountBits(tc.N)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for matrix %v", tc.N))
	}
}
