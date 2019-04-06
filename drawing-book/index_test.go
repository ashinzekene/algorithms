package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		NoOfPages int32
		Page      int32
		Expected  int32
	}{
		{
			40,
			4,
			2,
		},
		{
			6,
			2,
			1,
		},
		{
			5,
			4,
			0,
		},
	}
	for _, tc := range testCases {
		result := DrawingBook(tc.NoOfPages, tc.Page)
		assert.Equal(t, tc.Expected, result, "Failed for pages "+string(tc.Page))
	}
}
