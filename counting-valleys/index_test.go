package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		Path     string
		Steps    int32
		Expected int32
	}{
		{
			"UDDDUDUU",
			8,
			1,
		},
		{
			"DDUUDDUDUUUD",
			12,
			2,
		},
	}
	for _, tc := range testCases {
		result := CountingValleys(tc.Steps, tc.Path)
		assert.Equal(t, tc.Expected, result, "Failed for "+tc.Path)
	}
}
