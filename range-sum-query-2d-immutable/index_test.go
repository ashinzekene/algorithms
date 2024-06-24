package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Toys(t *testing.T) {
	testCases := []struct {
		Matrix     	[][]int
		Operations	[][4]int
		Expected 		[]int
	}{
		{
			[][]int{
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
			},
			[][4]int{
				{0, 0, 1, 1},
				{0, 0, 2, 2},
			},
			[]int{4, 9},
		},
	}
	for _, tc := range testCases {
		matrix := Constructor(tc.Matrix)
		for i, operation := range tc.Operations {
			assert.Equal(t, tc.Expected[i], matrix.SumRegion(operation[0], operation[1], operation[2], operation[3]))
		}
	}
}
