package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		K        int32
		Contests [][]int32
		Expected int32
	}{
		{
			3,
			[][]int32{
				[]int32{5, 1},
				[]int32{2, 1},
				[]int32{1, 1},
				[]int32{8, 1},
				[]int32{10, 0},
				[]int32{5, 0},
			},
			29,
		},
		{
			5,
			[][]int32{
				[]int32{13, 1},
				[]int32{10, 1},
				[]int32{9, 1},
				[]int32{8, 1},
				[]int32{13, 1},
				[]int32{12, 1},
				[]int32{18, 1},
				[]int32{13, 1},
			},
			42,
		},
		{
			2,
			[][]int32{
				[]int32{5, 1},
				[]int32{4, 0},
				[]int32{6, 1},
				[]int32{2, 1},
				[]int32{8, 0},
			},
			21,
		},
	}
	for _, tc := range testCases {
		actual := luckBalance(tc.K, tc.Contests)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Contests))
	}
}
