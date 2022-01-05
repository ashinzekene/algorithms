package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canFinish(t *testing.T) {
	testCases := []struct {
		NumCourses    int
		Prerequisites [][]int
		Expected      bool
	}{
		{
			2,
			[][]int{
				{1, 0},
			},
			true,
		},
		{
			2,
			[][]int{
				{1, 0},
				{0, 1},
			},
			false,
		},
		{
			5,
			[][]int{
				{1, 2},
				{4, 2},
				{1, 3},
			},
			true,
		},
		{
			5,
			[][]int{
				{1, 2},
				{4, 2},
				{2, 3},
				{3, 4},
			},
			false,
		},
		{
			7,
			[][]int{
				{1, 0},
				{0, 3},
				{0, 2},
				{3, 2},
				{2, 5},
				{4, 5},
				{5, 6},
				{2, 4},
			},
			true,
		},
		{
			7,
			[][]int{
				{1, 0},
				{0, 3},
				{0, 2},
				{3, 2},
				{2, 5},
				{4, 5},
				{5, 6},
				{2, 4},
				{5, 0},
			},
			false,
		},
	}
	for _, tc := range testCases {
		result := canFinish(tc.NumCourses, tc.Prerequisites)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("canFinish failed for matrix %d and %v", tc.NumCourses, tc.Prerequisites))
		result = canFinish1(tc.NumCourses, tc.Prerequisites)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("canFinish1 failed for matrix %d and %v", tc.NumCourses, tc.Prerequisites))
	}
}
