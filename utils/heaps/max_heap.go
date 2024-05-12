package heaps

import (
	"math"

	. "github.com/ashinzekene/algorithms/utils"
)

type MaxHeap struct {
	itemsCount int
	capacity int
	items []int
}


func NewMaxHeap(capacity int) MaxHeap {
	return MaxHeap{
		itemsCount: 0,
		capacity: capacity,
		items: make([]int, capacity),
	}
}

func (q *MaxHeap) GetItems() []int {
	return q.items
}


func (q *MaxHeap) Insert(val int) {
	if q.itemsCount == q.capacity {
		// panic
		return
	}
	q.items[q.itemsCount] = val
	q.itemsCount++
	q.heapifyFromBottom(q.itemsCount-1)
}

func (q *MaxHeap) Max() int {
	return q.items[0]
}

func (q *MaxHeap) RemoveMax() int {
	q.itemsCount--
	max_val := q.items[0]
	q.items[0] = q.items[q.itemsCount]
	q.items[q.itemsCount] = 0
	q.heapifyFromTop(0)
	return max_val
}

func (q *MaxHeap) GetLeftChildIndex(i int) int {
	return 2*i+1
}

func (q *MaxHeap) GetLeftChild(i int) int {
	index := q.GetLeftChildIndex(i)
	if index >= q.itemsCount {
		return int(math.Inf(-1))
	}
	return q.items[index]
}

func (q *MaxHeap) GetRightChildIndex(i int) int {
	return 2*i + 2
}	

func (q *MaxHeap) GetRightChild(i int) int {
	index := q.GetRightChildIndex(i)
	if index >= q.itemsCount {
		return int(math.Inf(-1))
	}
	return q.items[index]
}

func (q *MaxHeap) GetParentIndex(i int) int {
	return int(math.Round(float64(i) / 2)) - 1
}	

func (q *MaxHeap) GetParent(i int) int {
	index := q.GetParentIndex(i)
	if index < 0 {
		return int(math.Inf(1))
	}
	return q.items[index]
}

func (q *MaxHeap) heapifyFromTop(i int) {
	var max_val int = Max(q.items[i], q.GetLeftChild(i), q.GetRightChild(i))
	if q.items[i] == max_val {
		return
	}
	var new_i = q.GetRightChildIndex(i)
	if max_val == q.GetLeftChild(i) {
		new_i = q.GetLeftChildIndex(i)
	}
	var temp = q.items[i]
	q.items[i] = max_val
	q.items[new_i] = temp
	q.heapifyFromTop(new_i)
}

func (q *MaxHeap) heapifyFromBottom(i int) {
	if q.GetParentIndex(i) < 0 {
		return
	}
	var max_val = Max(q.GetParent(i), q.items[i])
	if max_val == q.GetParent(i) {
		return
	}
	var parentIndex = q.GetParentIndex(i)
	var temp = q.items[i]
	q.items[i] = q.GetParent(i)
	q.items[parentIndex] = temp
	q.heapifyFromBottom(parentIndex)
}

