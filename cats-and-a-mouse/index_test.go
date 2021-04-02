package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CatAndMouse(t *testing.T) {
	testCases := []struct {
		CatA     int32
		CatB     int32
		MouseC   int32
		Expected string
	}{
		{
			1,
			2,
			3,
			"Cat B",
		},
		{
			1,
			3,
			2,
			"Mouse C",
		},
	}
	for index, tc := range testCases {
		actual := CatsAndAMouse(tc.CatA, tc.CatB, tc.MouseC)
		assert.Equal(t, tc.Expected, actual, "Did not work for test case "+fmt.Sprint(index))
	}
}
