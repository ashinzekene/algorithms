package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_whatFlavors(t *testing.T) {
	testCases := []struct {
		Cost     []int32
		Money    int32
		Expected string
	}{
		{
			[]int32{1, 4, 5, 3, 2},
			4,
			"1 4",
		},
		{
			[]int32{2, 2, 4, 3},
			5,
			"1 4",
		},
		{
			[]int32{2, 2, 4, 3},
			4,
			"1 2",
		},
	}
	for _, tc := range testCases {
		actual := whatFlavors(tc.Cost, tc.Money)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Cost %d", tc.Cost))
	}
}
