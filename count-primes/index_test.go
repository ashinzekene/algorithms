package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NumTeams(t *testing.T) {
	testCases := []struct {
		N        int
		Expected int
	}{
		{
			10,
			4,
		},
	}
	for _, tc := range testCases {
		result := countPrimes(tc.N)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %d,", tc.N))
	}
}
