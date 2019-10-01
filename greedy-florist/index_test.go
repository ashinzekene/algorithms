package algorithms

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GreedyFlorist(t *testing.T) {
	testCases := []struct {
		FriendsCount int32
		Costs        []int32
		Expected     int32
	}{
		{
			3,
			[]int32{2, 5, 6},
			13,
		},
		{
			2,
			[]int32{2, 5, 6},
			15,
		},
		{
			3,
			[]int32{1, 3, 5, 7, 9},
			29,
		},
		{
			3,
			[]int32{390225, 426456, 688267, 800389, 990107, 439248, 240638, 15991, 874479, 568754, 729927, 980985, 132244, 488186, 5037, 721765, 251885, 28458, 23710, 281490, 30935, 897665, 768945, 337228, 533277, 959855, 927447, 941485, 24242, 684459, 312855, 716170, 512600, 608266, 779912, 950103, 211756, 665028, 642996, 262173, 789020, 932421, 390745, 433434, 350262, 463568, 668809, 305781, 815771, 550800},
			163578911,
		},
	}
	for _, tc := range testCases {
		actual := getMinimumCost(tc.FriendsCount, tc.Costs)
		assert.Equal(t, tc.Expected, actual, fmt.Sprintf("Failed for Digit %d", tc.Costs))
	}
}
