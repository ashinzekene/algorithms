package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		Squares  []int32
		Day      int32
		Month    int32
		Expected int32
		Message  string
	}{
		{
			[]int32{1, 2, 1, 3, 2},
			3,
			2,
			2,
			"Did not work with []int32{1, 2, 1, 3, 2}",
		},
		{
			[]int32{1, 1, 1, 1, 1},
			3,
			2,
			0,
			"Did not work with []int32{1, 1, 1, 1, 1}",
		},
		{
			[]int32{4},
			4,
			1,
			1,
			"Did not work with []int32{4}",
		},
		{
			[]int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1},
			18,
			7,
			3,
			"Did not work with []int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}",
		},
		{
			[]int32{3, 5, 4, 1, 2, 5, 3, 4, 3, 2, 1, 1, 2, 4, 2, 3, 4, 5, 3, 1, 2, 5, 4, 5, 4, 1, 1, 5, 3, 1, 4, 5, 2, 3, 2, 5, 2, 5, 2, 2, 1, 5, 3, 2, 5, 1, 2, 4, 3, 1, 5, 1, 3, 3, 5},
			18,
			6,
			10,
			"Did not work with []int32{3, 5, 4, 1, 2, 5, 3, 4, 3, 2, 1, 1, 2, 4, 2, 3, 4, 5, 3, 1, 2, 5, 4, 5, 4, 1, 1, 5, 3, 1, 4, 5, 2, 3, 2, 5, 2, 5, 2, 2, 1, 5, 3, 2, 5, 1, 2, 4, 3, 1, 5, 1, 3, 3, 5}",
		},
	}
	for _, tc := range testCases {
		result := BirthdayBar(tc.Squares, tc.Day, tc.Month)
		assert.Equal(t, tc.Expected, result, tc.Message)
	}
}
