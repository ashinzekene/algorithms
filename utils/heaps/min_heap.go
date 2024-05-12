package heaps

import (
	"math"

	. "github.com/ashinzekene/algorithms/utils"
)

type MinHeap struct {
	itemsCount int
	capacity int
	items []int
}


func NewMinHeap(capacity int) MinHeap {
	return MinHeap{
		itemsCount: 0,
		capacity: capacity,
		items: make([]int, capacity),
	}
}

func (q *MinHeap) GetItems() []int {
	return q.items
}


func (q *MinHeap) Insert(val int) {
	if q.itemsCount == q.capacity {
		// panic
		return
	}
	q.items[q.itemsCount] = val
	q.itemsCount++
	q.heapifyFromBottom(q.itemsCount-1)
}

func (q *MinHeap) Min() int {
	return q.items[0]
}

func (q *MinHeap) RemoveMin() {
	q.itemsCount--
	q.items[0] = q.items[q.itemsCount]
	q.items[q.itemsCount] = 0
	q.heapifyFromTop(0)
}

func (q *MinHeap) GetLeftChildIndex(i int) int {
	return 2*i+1
}

func (q *MinHeap) GetLeftChild(i int) int {
	index := q.GetLeftChildIndex(i)
	if index >= q.itemsCount {
		return int(math.Inf(1))
	}
	return q.items[index]
}

func (q *MinHeap) GetRightChildIndex(i int) int {
	return 2*i + 2
}	

func (q *MinHeap) GetRightChild(i int) int {
	index := q.GetRightChildIndex(i)
	if index >= q.itemsCount {
		return int(math.Inf(1))
	}
	return q.items[index]
}

func (q *MinHeap) GetParentIndex(i int) int {
	return int(math.Round(float64(i) / 2)) - 1
}	

func (q *MinHeap) GetParent(i int) int {
	index := q.GetParentIndex(i)
	if index < 0 {
		return int(math.Inf(-1))
	}
	return q.items[index]
}

func (q *MinHeap) heapifyFromTop(i int) {
	var min_val int = Min(q.items[i], q.GetLeftChild(i), q.GetRightChild(i))
	if q.items[i] == min_val {
		return
	}
	var new_i = q.GetRightChildIndex(i)
	if min_val == q.GetLeftChild(i) {
		new_i = q.GetLeftChildIndex(i)
	}
	var temp = q.items[i]
	q.items[i] = min_val
	q.items[new_i] = temp
	q.heapifyFromTop(new_i)
}

func (q *MinHeap) heapifyFromBottom(i int) {
	if q.GetParentIndex(i) < 0 {
		return
	}
	var min_val = Min(q.GetParent(i), q.items[i])
	if min_val == q.GetParent(i) {
		return
	}
	var parentIndex = q.GetParentIndex(i)
	var temp = q.items[i]
	q.items[i] = q.GetParent(i)
	q.items[parentIndex] = temp
	q.heapifyFromBottom(parentIndex)
}

