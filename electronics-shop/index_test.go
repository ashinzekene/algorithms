package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AngryProfessor(t *testing.T) {
	testCases := []struct {
		Keyboards []int32
		Drives    []int32
		Balance   int32
		Expected  int32
	}{
		{
			[]int32{3, 1},
			[]int32{5, 2, 8},
			10,
			9,
		},
		{
			[]int32{4},
			[]int32{5},
			5,
			-1,
		},
	}
	for _, tc := range testCases {
		actual := GetMoneySpent(tc.Keyboards, tc.Drives, tc.Balance)
		assert.Equal(t, tc.Expected, actual)
	}
}
