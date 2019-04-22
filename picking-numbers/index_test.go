package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PickingNumbers(t *testing.T) {
	testCases := []struct {
		Integers []int32
		Expected int32
	}{
		{
			[]int32{1, 1, 2, 2, 4, 4, 5, 5, 5},
			5,
		},
		{
			[]int32{4, 6, 5, 3, 3, 1},
			3,
		},
		{
			[]int32{1, 2, 2, 3, 1, 2},
			5,
		},
	}
	for _, tc := range testCases {
		actual := PickingNumbers(tc.Integers)
		assert.Equal(t, tc.Expected, actual)
	}
}
