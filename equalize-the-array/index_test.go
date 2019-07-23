package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EqualizeArray(t *testing.T) {
	testCases := []struct {
		Arr      []int32
		Expected int32
	}{
		{
			[]int32{3, 3, 2, 1, 3},
			2,
		},
		{
			[]int32{1, 2, 2, 3},
			2,
		},
		{
			[]int32{1, 2, 2, 3, 4, 4, 4, 5},
			5,
		},
		{
			[]int32{4, 7, 8, 3, 4, 4, 4, 5, 1, 8, 8, 9, 8, 8},
			9,
		},
	}
	for _, tc := range testCases {
		actual := EqualizeArray(tc.Arr)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for List %d", tc.Arr))
	}
}
