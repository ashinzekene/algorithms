package algorithms

type Queue[T any] struct {
	list []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.list = append(q.list, val)
}

func (q *Queue[T]) Dequeue() T {
	// Queue might be empty
	val := q.list[0]
	q.list = q.list[1:len(q.list)]
	return val
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue[T]) Peek() T {
	return q.list[0]
}
