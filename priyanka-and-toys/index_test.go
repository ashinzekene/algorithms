package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Toys(t *testing.T) {
	testCases := []struct {
		Toys     []int32
		Expected int32
	}{
		{
			[]int32{1, 2, 3, 21, 7, 12, 14, 21},
			4,
		},
		{
			[]int32{12, 15, 7, 8, 19, 24},
			4,
		},
	}
	for _, tc := range testCases {
		actual := toys(tc.Toys)
		assert.Equal(t, tc.Expected, actual)
	}
}
