package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reorderLogFiles(t *testing.T) {
	testCases := []struct {
		Log      []string
		Expected []string
	}{
		{
			[]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo", "a2 act car"},
			[]string{"a2 act car", "g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"},
		},
		{
			[]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"},
			[]string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"},
		},
	}
	for _, tc := range testCases {
		actual := reorderLogFiles(tc.Log)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Log))
	}
}
