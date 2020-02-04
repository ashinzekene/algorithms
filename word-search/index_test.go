package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_WordSearch(t *testing.T) {
	board := [][]byte{
		[]byte{'A', 'B', 'C', 'E'},
		[]byte{'S', 'F', 'C', 'S'},
		[]byte{'A', 'D', 'E', 'E'},
	}
	testCases := []struct {
		Word     string
		Expected bool
	}{
		{
			"SEE",
			true,
		},
		{
			"SAW",
			false,
		},
		{
			"SEED",
			true,
		},
		{
			"SEEDASFCS",
			false,
		},
		{
			"S",
			true,
		},
		{
			"SAD",
			true,
		},
	}
	for _, tc := range testCases {
		actual := exist(board, tc.Word)
		assert.Equal(t, tc.Expected, actual, "Did not work for "+tc.Word)
	}
}
