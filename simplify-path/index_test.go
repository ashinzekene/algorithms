package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_SimplifyPath(t *testing.T) {
	testCases := []struct {
		Path     string
		Expected string
	}{
		{
			"/a/./b/..///../c/",
			"/c",
		},
		{
			"/home//foo/",
			"/home/foo",
		},
		{
			"/../",
			"/",
		},
		{
			"/...",
			"/...",
		},
		{
			"/a//b////c/d//././/..",
			"/a/b/c",
		},
	}
	for _, tc := range testCases {
		actual := SimplifyPath(tc.Path)
		assert.Equal(t, tc.Expected, actual)
	}
}
