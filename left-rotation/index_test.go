package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		Rotations int32
		Array     []int32
		Expected  []int32
	}{
		{
			4,
			[]int32{1, 2, 3, 4, 5},
			[]int32{5, 1, 2, 3, 4},
		},
		{
			33,
			[]int32{6, 2, 3, 4, 5, 7, 9},
			[]int32{7, 9, 6, 2, 3, 4, 5},
		},
	}
	for _, tc := range testCases {
		actual := LeftRotation(tc.Array, tc.Rotations)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not wortk for %v", tc.Array))
	}
}
