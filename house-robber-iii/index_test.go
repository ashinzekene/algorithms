package algorithms

import (
	"fmt"
	"testing"

	. "github.com/ashinzekene/algorithms/utils/trees"
	"github.com/stretchr/testify/assert"
)

func Test_rob(t *testing.T) {
	testCases := []struct {
		Root     *Tree
		Expected int
	}{
		{
			TreeFromListLevelOrder([]int{3, 2, 3, -1, 3, -1, 1}, -1),
			7,
		},
		{
			TreeFromListLevelOrder([]int{3, 4, 5, 1, 3, -1, 1}, -1),
			9,
		},
	}
	for _, tc := range testCases {
		actual := rob(tc.Root.Head)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for %v", tc.Root.LevelOrder()))
	}
}
