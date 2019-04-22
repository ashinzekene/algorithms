package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		Ad       int32
		Expected int32
	}{
		{
			0,
			0,
		},
		{
			2,
			1,
		},
		{
			5,
			2,
		},
		{
			24,
			5,
		},
	}
	for _, tc := range testCases {
		actual := viralAvertising(tc.Ad)
		assert.Equal(t, tc.Expected, actual)
	}
}
