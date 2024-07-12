package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findCheapestPrice(t *testing.T) {
	testCases := []struct {
		N        int
		Flights  [][]int
		Src      int
		Dest     int
		K        int
		Expected int
	}{
		{
			4,
			[][]int{
				{0, 1, 100},
				{1, 2, 100},
				{2, 0, 100},
				{1, 3, 600},
				{2, 3, 200},
			},
			0,
			3,
			1,
			700,
		},
		{
			10,
			[][]int{
				{3, 4, 4},
				{2, 5, 6},
				{4, 7, 10},
				{9, 6, 5},
				{7, 4, 4},
				{6, 2, 10},
				{6, 8, 6},
				{7, 9, 4},
				{1, 5, 4},
				{1, 0, 4},
				{9, 7, 3},
				{7, 0, 5},
				{6, 5, 8},
				{1, 7, 6},
				{4, 0, 9},
				{5, 9, 1},
				{8, 7, 3},
				{1, 2, 6},
				{4, 1, 5},
				{5, 2, 4},
				{1, 9, 1},
				{7, 8, 10},
				{0, 4, 2},
				{7, 2, 8},
			},
			6,
			0,
			7,
			14,
		},
		{
			3,
			[][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			0,
			2,
			1,
			200,
		},
		{
			3,
			[][]int{
				{0, 1, 100},
				{1, 2, 100},
				{0, 2, 500},
			},
			0,
			2,
			0,
			500,
		},
	}
	for index, tc := range testCases {
		actual := findCheapestPrice(tc.N, tc.Flights, tc.Src, tc.Dest, tc.K)
		assert.Equal(t, tc.Expected, actual, "Did not work for test case "+ fmt.Sprint(index))
	}
}
