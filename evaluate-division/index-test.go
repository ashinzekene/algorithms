package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calcEquation(t *testing.T) {
	testCases := []struct {
		Equation    [][]string
		Values			[]float64
		Queries 		[][]string
		Expected	 	[]float64
	}{
		{
			[][]string{
				{"a","b"},
				{"b","c"},
			},
			[]float64{2.0,3.0},
			[][]string{
				{"a","c"},
				{"b","a"},
				{"a","e"},
				{"a","a"},
				{"x","x"},
			},
			[]float64{6, 0.5, -1, 1, -1},
		},
		{
			[][]string{
				{"a","b"},
				{"b","c"},
				{"bc","cd"},
			},
			[]float64{1.5,2.5,5.0},
			[][]string{
				{"a","c"},
				{"c","b"},
				{"bc","cd"},
				{"cd","bc"},
			},
			[]float64{3.75, 0.4, 5, 0.2},
		},
		{
				[][]string{
				{"a","b"},
			},
			[]float64{0.5},
			[][]string{
				{"a","b"},
				{"b","a"},
				{"a","c"},
				{"x","y"},
			},
			[]float64{0.5, 2, -1, -1},
		},
	}
	for _, tc := range testCases {
		actual := calcEquation(tc.Equation, tc.Values, tc.Queries)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Equation %v", tc.Equation))
	}
}
