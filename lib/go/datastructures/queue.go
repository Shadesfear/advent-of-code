package datastructures

type Queue[T any] struct {
	Items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(item T) {
	q.Items = append(q.Items, item)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.Items) == 0 {
		var zero T
		return zero, false
	}
	item := q.Items[0]
	q.Items = q.Items[1:]
	return item, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if len(q.Items) == 0 {
		var zero T
		return zero, false
	}
	item := q.Items[0]
	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.Items) == 0
}

func (q *Queue[T]) Len() int {
	return len(q.Items)
}
