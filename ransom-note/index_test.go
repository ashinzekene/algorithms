package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RandomNote(t *testing.T) {
	testCases := []struct {
		Magazine []string
		Note     []string
		Expected string
	}{
		{
			[]string{"give", "me", "one", "grand", "today", "night"},
			[]string{"give", "one", "grand", "today"},
			"Yes",
		},
		{
			[]string{"two", "times", "three", "is", "not", "four"},
			[]string{"two", "times", "two", "is", "four"},
			"No",
		},
		{
			[]string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"},
			[]string{"ive", "got", "some", "coconuts"},
			"No",
		},
	}
	for _, tc := range testCases {
		actual := checkMagazine(tc.Magazine, tc.Note)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Did not work for %v", tc.Note))
	}
}
