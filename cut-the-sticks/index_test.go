package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CutTheSticks(t *testing.T) {
	testCases := []struct {
		Arr      []int32
		Expected []int32
	}{
		{
			[]int32{1, 2, 3, 4, 3, 3, 2, 1},
			[]int32{8, 6, 4, 1},
		},
		{
			[]int32{5, 4, 4, 2, 2, 8},
			[]int32{6, 4, 2, 1},
		},
		{
			[]int32{4, 4, 4, 2, 2, 8},
			[]int32{6, 4, 1},
		},
	}
	for _, tc := range testCases {
		result := CutTheSticks(tc.Arr)
		assert.Equal(t, tc.Expected, result, "Failed for %v", tc.Arr)
	}
}
