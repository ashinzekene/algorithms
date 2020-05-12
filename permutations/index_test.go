package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Permutations(t *testing.T) {
	testCases := []struct {
		Nums        int
		Expected      [][]int
	}{
		{
			[]int{1, 2}
			[][]int{
				[]int{1, 2},
				[]int{2, 1},
			}
		},
		{
			[]int{1, 2, 3}
			[][]int{
				[]int{1, 2, 3},
				[]int{1, 3, 2},
				[]int{2, 1, 3},
				[]int{2, 3, 1},
				[]int{3, 1, 2},
				[]int{3, 2, 1},
			}
		},
		{
			[]int{1, 2, 3, 4}
			[][]int{
				[]int{1, 2, 3, 4},
				[]int{1, 2, 4, 3},
				[]int{1, 3, 2, 4},
				[]int{1, 3, 4, 2},
				[]int{1, 4, 2, 3},
				[]int{1, 4, 3, 2},
				[]int{2, 1, 3, 4},
				[]int{2, 1, 4, 3},
				[]int{2, 3, 1, 4},
				[]int{2, 3, 4, 1},
				[]int{2, 4, 1, 3},
				[]int{2, 4, 3, 1},
				[]int{3, 1, 2, 4},
				[]int{3, 1, 4, 1},
				[]int{3, 2, 1, 4},
				[]int{3, 2, 4, 1},
				[]int{3, 4, 1, 2},
				[]int{3, 4, 2, 1},
				[]int{4, 1, 2, 3},
				[]int{4, 1, 3, 1},
				[]int{4, 2, 1, 3},
				[]int{4, 2, 3, 1},
				[]int{4, 3, 1, 2},
				[]int{4, 3, 2, 1},
			},
		},
		{
			[][]int{
				[]int{2, 6, 8},
				[]int{3, 5, 7},
				[]int{1, 8, 1},
				[]int{5, 9, 15},
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
