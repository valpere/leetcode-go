package stack

import (
	"errors"
	"fmt"
)

type StackInterface[T any] interface {
	Size() int
	Len() int
	IsEmpty() bool
	IsFull() bool
	Push(elem T) error
	Peek() (elem T, err error)
	Pop() (elem T, err error)
}

type Stack[T any] struct {
	elements []T
	topPos   int
}

var ErrStackIsEmpty = errors.New("is empty")

func NewStack[T any](size int) *Stack[T] {
	if size < 0 {
		size = 0
	}

	return &Stack[T]{elements: make([]T, size), topPos: -1}
}

func (q *Stack[T]) Top() int {
	return q.topPos
}

func (q *Stack[T]) IsEmpty() bool {
	return (q.topPos < 0)
}

func (q *Stack[T]) Push(elem T) error {
	if q.topPos < len(q.elements)-1 {
		q.elements[q.topPos] = elem
	} else {
		q.elements = append(q.elements, elem)
	}

	q.topPos++
	return nil
}

func (q *Stack[T]) Peek() (elem T, err error) {
	if q.IsEmpty() {
		return elem, ErrStackIsEmpty
	}

	return q.elements[q.topPos], nil
}

func (q *Stack[T]) Pop() (elem T, err error) {
	if q.IsEmpty() {
		return elem, ErrStackIsEmpty
	}

	elem = q.elements[q.topPos]
	q.topPos--

	return elem, nil
}

func (q Stack[T]) String() string {
	return fmt.Sprintf("elements: %v; top: %v", q.elements, q.topPos)
}

func SwapSkacks[T any](ist, ost *Stack[T]) {
	for !ist.IsEmpty() {
		elem, err := ist.Pop()
		if err != nil {
			panic(err)
		}
		ost.Push(elem)
	}
}
