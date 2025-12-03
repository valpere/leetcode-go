package queue

import "errors"

type QueueInterface[T any] interface {
	Size() int
	IsEmpty() bool
	IsFull() bool
	Enqueue(elem T) error
	Peek() (elem T, err error)
	Dequeue() (elem T, err error)
}

type Queue[T any] struct {
	elements []T
	outPos   int
	inpPos   int
	count    int
}

func NewQueue[T any](size int) *Queue[T] {
	q := new(Queue[T])
	if size == 0 {
		size = 8
	}
	q.elements = make([]T, size)
	return q
}

func (q *Queue[T]) Size() int {
	return len(q.elements)
}

func (q *Queue[T]) Count() int {
	return q.count
}

func (q *Queue[T]) IsEmpty() bool {
	return (q.Count() == 0)
}

func (q *Queue[T]) IsFull() bool {
	return (q.Count() == q.Size())
}

func (q *Queue[T]) Enqueue(elem T) error {
	if q.IsFull() {
		return errors.New("is full")
	}

	q.elements[q.inpPos] = elem

	q.count++
	q.inpPos++
	if q.inpPos >= q.Size() {
		q.inpPos = 0
	}

	return nil
}

func (q *Queue[T]) Peek() (elem T, err error) {
	if q.IsEmpty() {
		return elem, errors.New("is empty")
	}

	return q.elements[q.outPos], nil
}

func (q *Queue[T]) Dequeue() (elem T, err error) {
	if q.IsEmpty() {
		return elem, errors.New("is empty")
	}

	elem = q.elements[q.outPos]
	q.count--
	q.outPos++
	if q.outPos >= q.Size() {
		q.outPos = 0
	}

	return elem, nil
}
