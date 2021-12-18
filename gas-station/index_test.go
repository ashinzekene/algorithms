package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CanCompleteCircuit(t *testing.T) {
	testCases := []struct {
		Gas      []int
		Cost     []int
		Expected int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			[]int{3, 4, 5, 1, 2},
			3,
		},
		{
			[]int{2, 3, 4},
			[]int{3, 4, 3},
			-1,
		},
		{
			[]int{2},
			[]int{3},
			-1,
		},
		{
			[]int{2},
			[]int{2},
			0,
		},
		{
			[]int{5, 1, 2, 3, 4},
			[]int{4, 4, 1, 5, 1},
			4,
		},
	}
	for _, tc := range testCases {
		actual := CanCompleteCircuit(tc.Gas, tc.Cost)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for Gas: %v and Cost: %v", tc.Gas, tc.Cost))
	}
}
