package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WordSearch(t *testing.T) {
	testCases := []struct {
		Board    [][]byte
		Word     string
		Expected bool
	}{
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEE",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SAW",
			false,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEED",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SEEDASFCS",
			false,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"S",
			true,
		},
		{
			[][]byte{
				{'A', 'B', 'C', 'E'},
				{'S', 'F', 'C', 'S'},
				{'A', 'D', 'E', 'E'},
			},
			"SAD",
			true,
		},
		{
			[][]byte{
				{'A', 'A'},
			},
			"AAA",
			false,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'C', 'D'},
			},
			"ABCD",
			false,
		},
		{
			[][]byte{
				{'A', 'B'},
				{'D', 'C'},
			},
			"ABCD",
			true,
		},
	}
	for _, tc := range testCases {
		actual := exist(tc.Board, tc.Word)
		assert.Equal(t, tc.Expected, actual, "Did not work for "+tc.Word)
	}
}
