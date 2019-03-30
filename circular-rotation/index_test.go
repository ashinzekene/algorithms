package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CircularRotation(t *testing.T) {
	cases := []struct {
		Expected  []int32
		Array     []int32
		Queries   []int32
		Rotations int32
		Message   string
	}{
		{
			[]int32{2, 3, 1},
			[]int32{1, 2, 3},
			[]int32{0, 1, 2},
			2,
			"Expected [2,3,1] to work",
		},
	}

	for _, tc := range cases {
		actual := CircularArrayRotation(tc.Array, tc.Rotations, tc.Queries)
		assert.Equal(t, tc.Expected, actual, tc.Message)
	}
}
