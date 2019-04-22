package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ViralAdvertising(t *testing.T) {
	testCases := []struct {
		Ad       int32
		Expected int32
	}{
		{
			0,
			0,
		},
		{
			5,
			24,
		},
		{
			4,
			15,
		},
		{
			3,
			9,
		},
	}
	for _, tc := range testCases {
		actual := ViralAdvertising(tc.Ad)
		assert.Equal(t, tc.Expected, actual)
	}
}
