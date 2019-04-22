package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MergeSort(t *testing.T) {
	testCases := []struct {
		Numbers  []int
		Expected []int
	}{
		{
			[]int{1, 4, 5, 2, 6, 7, 9, 8, 3, 24, 72, 14, 13, 55, 42, 31, 30, 29, 28, 27, 20},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 13, 14, 20, 24, 27, 28, 29, 30, 31, 42, 55, 72},
		},
	}
	for _, tc := range testCases {
		actual := MergeSort(tc.Numbers)
		assert.Equal(t, tc.Expected, actual)
	}
}
