package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MaximalRectangle(t *testing.T) {
	testCases := []struct {
		Matrix   [][]byte
		Expected int
	}{
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			6,
		},
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '0', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			},
			5,
		},
		{
			[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '1', '1', '1'},
			},
			9,
		},
		{
			[][]byte{
				{'1', '0', '1', '0'},
				{'1', '0', '1', '1'},
				{'1', '1', '1', '1'},
			},
			4,
		},
		{
			[][]byte{
				{'1', '1'},
				{'1', '0'},
			},
			2,
		},
		{
			[][]byte{
				{'1'},
			},
			1,
		},
		{
			[][]byte{
				{'0'},
			},
			0,
		},
		{
			[][]byte{},
			0,
		},
	}
	for _, tc := range testCases {
		actual := MaximalRectangle(tc.Matrix)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Matrix))
	}
}
