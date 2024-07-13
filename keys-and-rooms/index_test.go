package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canVisitAllRooms(t *testing.T) {
	testCases := []struct {
		Paths    [][]int
		Expected bool
	}{
		{
			[][]int{
				{1},
				{2},
				{3},
				{},
			},
			true,
		},
		{
			[][]int{
				{1,3},
				{3,0,1},
				{2},
				{0},
			},
			false,
		},
	}
	for index, tc := range testCases {
		actual := canVisitAllRooms(tc.Paths)
		assert.Equal(t, tc.Expected, actual, "Did not work for test case "+fmt.Sprint(index))
	}
}
