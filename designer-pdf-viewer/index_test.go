package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DesignerPDF(t *testing.T) {
	testCases := []struct {
		Height   []int32
		Word     string
		Expected int32
	}{
		{
			[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
			"abc",
			9,
		},
		{
			[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
			"zaba",
			28,
		},
		{
			[]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7},
			"guard",
			25,
		},
	}
	for _, tc := range testCases {
		result := DesignerPdfViewer(tc.Height, tc.Word)
		assert.Equal(t, tc.Expected, result, "Failed for "+tc.Word)
	}
}
