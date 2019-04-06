package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AngryProfessor(t *testing.T) {
	testCases := []struct {
		Array    []int32
		Expected []int32
	}{
		{
			[]int32{4, 3, 5, 1, 2},
			[]int32{1, 3, 5, 4, 2},
		},
		{
			[]int32{2, 3, 1},
			[]int32{2, 3, 1},
		},
	}
	for _, tc := range testCases {
		actual := PermutationEquation(tc.Array)
		assert.Equal(t, tc.Expected, actual)
	}
}
