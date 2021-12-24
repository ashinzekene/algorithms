package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solve(t *testing.T) {
	testCases := []struct {
		Board    [][]byte
		Expected [][]byte
	}{
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'O', 'X'},
			},
		},
		{
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'X', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
			},
		},
		{
			[][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'O', 'X'},
				{'X', 'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X', 'X'},
			},
			[][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X', 'X'},
			},
		},
		{
			[][]byte{
				{'X'},
			},
			[][]byte{
				{'X'},
			},
		},
	}
	for _, tc := range testCases {
		solve(tc.Board)
		assert.ElementsMatch(t, tc.Expected, tc.Board, fmt.Sprintf("Failed for %v", tc.Board))
	}
}
