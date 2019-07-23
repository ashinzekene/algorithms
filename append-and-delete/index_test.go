package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AppendAndDelete(t *testing.T) {
	testCases := []struct {
		S        string
		T        string
		K        int32
		Expected string
	}{
		{
			"hackerhappy",
			"hackerrank",
			9,
			"Yes",
		},
		{
			"aba",
			"aba",
			7,
			"Yes",
		},
		{
			"aaaaaaaaaa",
			"aaaaa",
			7,
			"Yes",
		},
		{
			"zzzzz",
			"zzzzzzz",
			4,
			"Yes",
		},
		{
			"y",
			"yu",
			2,
			"No",
		},
		{
			"abcd",
			"abcdert",
			10,
			"No",
		},
		{
			"aaa",
			"a",
			5,
			"Yes",
		},
		{
			"ashley",
			"ash",
			2,
			"No",
		},
		{
			"asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv",
			"asdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcvasdfqwertyuighjkzxcv",
			20,
			"Yes",
		},
	}
	for i, tc := range testCases {
		result := AppendAndDelete(tc.S, tc.T, tc.K)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v", i))
	}
}
