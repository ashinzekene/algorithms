package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PickingNumbers(t *testing.T) {
	testCases := []struct {
		Integers []int
		Expected int
	}{
		{
			[]int{1, 1, 2, 2, 4, 4, 5, 5, 5},
			5,
		},
		{
			[]int{4, 6, 5, 3, 3, 1},
			3,
		},
		{
			[]int{1, 2, 2, 3, 1, 2},
			5,
		},
	}
	for _, tc := range testCases {
		actual := PickingNumbers(tc.Integers)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("PickingNumbers did not work for %v", tc.Integers))
		actual = PickingNumbers1(tc.Integers)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("PickingNumbers1 did not work for %v", tc.Integers))
	}
}
