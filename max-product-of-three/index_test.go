package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaxProductThree(t *testing.T) {
	testCases := []struct {
		Array    []int
		Expected int
	}{
		{
			[]int{-7,0,1,2,5,6},
			60,
		},
		{
			[]int{-7,-2,1,2,5,6},
			84,
		},
		{
			[]int{-7,-1,1,2,5,6},
			60,
		},
		{
			[]int{1,2,5,6},
			60,
		},
	}
	for _, tc := range testCases {
		actual := maxProductThree(tc.Array)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Array))
	}
}
