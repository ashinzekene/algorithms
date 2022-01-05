package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numIslands(t *testing.T) {
	testCases := []struct {
		Grid     [][]byte
		Expected int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
	}
	for _, tc := range testCases {
		actual := numIslands(tc.Grid)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Grid))
	}
}
