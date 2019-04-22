package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AngryProfessor(t *testing.T) {
	testCases := []struct {
		Tree     int32
		Expected int32
	}{
		{
			0,
			1,
		},
		{
			1,
			2,
		},
		{
			4,
			7,
		},
		{
			5,
			14,
		},
	}
	for _, tc := range testCases {
		actual := UtopianTree(tc.Tree)
		assert.Equal(t, tc.Expected, actual)
	}
}
