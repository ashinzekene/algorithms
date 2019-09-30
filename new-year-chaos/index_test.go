package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewYearChaos(t *testing.T) {
	testCases := []struct {
		Array    []int32
		Expected interface{}
	}{
		{
			[]int32{2, 1, 5, 3, 4},
			3,
		},
		{
			[]int32{2, 5, 1, 3, 4},
			"Too chaotic",
		},
	}
	for _, tc := range testCases {
		actual := MinimumBribes(tc.Array)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Array))
	}
}
