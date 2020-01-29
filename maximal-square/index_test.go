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
				[]byte{49, 49, 48, 49},
				[]byte{49, 49, 48, 49},
				[]byte{49, 49, 49, 49},
			},
			4,
		},
		{
			[][]byte{
				[]byte{49, 48, 49, 48, 48},
				[]byte{49, 48, 49, 49, 49},
				[]byte{49, 49, 49, 49, 49},
				[]byte{49, 48, 48, 49, 48},
			},
			4,
		},
	}
	for i, tc := range testCases {
		result := maximalSquare(tc.Arr)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed at index %v", i))
	}
}
