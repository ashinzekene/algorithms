package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Contacts(t *testing.T) {
	testCases := []struct {
		Queries  [][]string
		Expected []int32
	}{
		{
			[][]string{
				{"add", "ed"},
				{"add", "eddie"},
				{"add", "edward"},
				{"find", "ed"},
				{"add", "edwina"},
				{"find", "e"},
				{"find", "a"},
			},
			[]int32{3, 4, 0},
		},
		{
			[][]string{
				{"add", "hack"},
				{"add", "hackerrank"},
				{"find", "hac"},
				{"find", "hak"},
			},
			[]int32{2, 0},
		},
		{
			[][]string{
				{"add", "ed"},
				{"add", "eddie"},
				{"add", "edward"},
				{"add", "edmund"},
				{"add", "elephant"},
				{"find", "e"},
				{"find", "ed"},
				{"find", "edw"},
				{"find", "a"},
			},
			[]int32{5, 4, 1, 0},
		},
	}
	for _, tc := range testCases {
		result := Contacts(tc.Queries)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v ", tc.Queries))
	}
}
