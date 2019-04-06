package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AngryProfessor(t *testing.T) {
	testCases := []struct {
		Threshold int32
		Array     []int32
		Expected  string
		Message   string
	}{
		{
			4,
			[]int32{-93, -86, 49, -62, -90, -63, 40, 72, 11, 67},
			"NO",
			"Did not work",
		},
		{
			10,
			[]int32{23, -35, -2, 58, -67, -56, -42, -73, -19, 37},
			"YES",
			"Did not work",
		},
		{
			9,
			[]int32{13, 91, 56, -62, 96, -5, -84, -36, -46, -13},
			"YES",
			"Did not work",
		},
		{
			8,
			[]int32{45, 67, 64, -50, -8, 78, 84, -51, 99, 60},
			"YES",
			"Did not work",
		},
		{
			7,
			[]int32{26, 94, -95, 34, 67, -97, 17, 52, 1, 86},
			"YES",
			"Did not work",
		},
		{
			2,
			[]int32{19, 29, -17, -71, -75, -27, -56, -53, 65, 83},
			"NO",
			"Did not work",
		},
		{
			10,
			[]int32{-32, -3, -70, 8, -40, -96, -18, -46, -21, -79},
			"YES",
			"Did not work",
		},
		{
			9,
			[]int32{-50, 0, 64, 14, -56, -91, -65, -36, 51, -28},
			"YES",
			"Did not work",
		},
		{
			3,
			[]int32{-1, -3, 4, 2},
			"YES",
			"Did not work",
		},
		{
			2,
			[]int32{-1, -3, 4, 2},
			"NO",
			"Did not work",
		},
		{
			5,
			[]int32{-1, -3, 4, 2, -3, 1},
			"YES",
			"Did not work",
		},
	}
	for _, tc := range testCases {
		actual := AngryProfessor(tc.Threshold, tc.Array)
		assert.Equal(t, tc.Expected, actual, tc.Message)
	}
}
