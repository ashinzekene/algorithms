package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_allPathsSourceTarget(t *testing.T) {
	testCases := []struct {
		Paths    [][]int
		Expected [][]int
	}{
		{
			[][]int{
				{1, 2},
				{3},
				{3},
				{},
			},
			[][]int{
				{0, 1, 3},
				{0, 2, 3},
			},
		},
		{
			[][]int{
				{4, 3, 1},
				{3, 2, 4},
				{3},
				{4},
				{},
			},
			[][]int{
				{0, 4},
				{0, 3, 4},
				{0, 1, 3, 4},
				{0, 1, 2, 3, 4},
				{0, 1, 4},
			},
		},
		{
			[][]int{
				{3, 1},
				{4, 6, 7, 2, 5},
				{4, 6, 3},
				{6, 4},
				{7, 6, 5},
				{6},
				{7},
				{},
			},
			[][]int{
				{0, 3, 6, 7},
				{0, 3, 4, 7},
				{0, 3, 4, 6, 7},
				{0, 3, 4, 5, 6, 7},
				{0, 1, 4, 7},
				{0, 1, 4, 6, 7},
				{0, 1, 4, 5, 6, 7},
				{0, 1, 6, 7},
				{0, 1, 7},
				{0, 1, 2, 4, 7},
				{0, 1, 2, 4, 6, 7},
				{0, 1, 2, 4, 5, 6, 7},
				{0, 1, 2, 6, 7},
				{0, 1, 2, 3, 6, 7},
				{0, 1, 2, 3, 4, 7},
				{0, 1, 2, 3, 4, 6, 7},
				{0, 1, 2, 3, 4, 5, 6, 7},
				{0, 1, 5, 6, 7},
			},
		},
	}
	for index, tc := range testCases {
		actual := allPathsSourceTarget(tc.Paths)
		assert.Equal(t, tc.Expected, actual, "Did not work for test case "+fmt.Sprint(index))
	}
}
