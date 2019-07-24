
package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Squares(t *testing.T) {
	testCases := []struct{
		A int32
		B int32
		Expected int32
	}{
		{
			3,
			9,
			2,
		},
		{
			17,
			24,
			0,
		},
		{
			35,
			70,
			3,
		},
		{
			100,
			1000,
			22,
		},
		{
			50,
			1050,
			25,
		},
	}
	for _, tc := range testCases {
		result := Squares(tc.A, tc.B)
		assert.Equal(t, tc.Expected, result, fmt.Sprintf("Failed for %v and %v", tc.A , tc.B))
	}
}