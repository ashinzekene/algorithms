package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ArrayManipulation(t *testing.T) {
	testCases := []struct {
		Arr      [][]int32
		N        int32
		Expected int64
	}{
		{
			[][]int32{
				[]int32{1, 2, 100},
				[]int32{2, 5, 100},
				[]int32{3, 4, 100},
			},
			5,
			200,
		},
		{
			[][]int32{
				[]int32{1, 5, 3},
				[]int32{4, 8, 7},
				[]int32{6, 9, 1},
			},
			10,
			10,
		},
		{
			[][]int32{
				[]int32{2, 6, 8},
				[]int32{3, 5, 7},
				[]int32{1, 8, 1},
				[]int32{5, 9, 15},
			},
			10,
			31,
		},
	}
	for i, tc := range testCases {
		result := ArrayManipulation(tc.N, tc.Arr)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", i))
	}
}
