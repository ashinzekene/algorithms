package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CountSortedVowelStrings(t *testing.T) {
	testCases := []struct {
		Count    int
		Expected int
	}{
		{
			1,
			5,
		},
		{
			2,
			15,
		},
		{
			3,
			35,
		},
	}
	for _, tc := range testCases {
		result := countVowelStrings(tc.Count)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", tc.Count))
	}
}
