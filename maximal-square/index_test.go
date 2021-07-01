package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximalSquare(t *testing.T) {
	testCases := []struct {
		Arr      [][]byte
		Expected int
	}{
		{
			[][]byte{
				{'1', '1', '0', '1'},
				{'1', '1', '0', '1'},
				{'1', '1', '1', '1'},
			},
			4,
		},
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			4,
		},
		{
			[][]byte{
				{'0', '1'},
				{'1', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			9,
		},
	}
	for i, tc := range testCases {
		result := maximalSquare(tc.Arr)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed at index %v", i))
	}
}
