package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_LeftRotation(t *testing.T) {
	testCases := []struct {
		Count    int32
		Array    []int32
		Expected int32
	}{
		{
			4,
			[]int32{1, 2, 3, 4, 10, 20, 30, 40, 100, 200},
			3,
		},
		{
			2,
			[]int32{1, 2, 1, 2, 1},
			0,
		},
		{
			3,
			[]int32{10, 100, 300, 200, 1000, 20, 30},
			20,
		},
	}
	for _, tc := range testCases {
		actual := maxMin(tc.Count, tc.Array)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Array))
	}
}
