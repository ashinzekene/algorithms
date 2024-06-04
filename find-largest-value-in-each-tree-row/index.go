package algorithms

import (
	"math"

	"github.com/ashinzekene/algorithms/utils"
	"github.com/ashinzekene/algorithms/utils/queues"
	"github.com/ashinzekene/algorithms/utils/trees"
)

func largestValues(root *trees.TreeNode) []int {
	type ValRow struct {
		val *trees.TreeNode
		row int
	}
	

	queue := queues.NewQueue[ValRow]()
	queue.Enqueue(ValRow{
		val: root,
		row: 0,
	})

	maxes := make([]int, 0)
	curr_index := 0
	current_max := math.Inf(-1)
	for !queue.IsEmpty() {
		curr :=  queue.Dequeue()
		if curr.val.Left != nil {
			queue.Enqueue(ValRow{
				val: curr.val.Left,
				row: curr.row + 1,
			})
		}
		if curr.val.Right != nil {
			queue.Enqueue(ValRow{
				val: curr.val.Right,
				row: curr.row + 1,
			})
		}
		if curr.row != curr_index {
			maxes = append(maxes, int(current_max))
			current_max = float64(curr.val.Val)
			curr_index++
		} else if curr.val != nil {
			current_max = utils.Max(current_max, float64(curr.val.Val))
		}
	}

	maxes = append(maxes, int(current_max))

	return maxes
}
