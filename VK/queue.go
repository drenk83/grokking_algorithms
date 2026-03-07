package main

import "errors"

var ErrFull = errors.New("queue is full")
var ErrEmpty = errors.New("queue is empty")

type Queue struct {
	items []string
	size  int
	head  int
}

func (q *Queue) Push(elem string) error {
	if q.size == len(q.items) {
		return ErrFull
	}

	tail := q.head + q.size
	if tail >= len(q.items) {
		tail -= len(q.items)
	}
	q.items[tail] = elem
	q.size++
	return nil
}

func (q *Queue) Pop() (string, error) {
	if q.size == 0 {
		return "", ErrEmpty
	}

	out := q.items[q.head]
	q.head++
	if q.head == len(q.items) {
		q.head = 0
	}

	q.size--
	return out, nil
}

func NewQueue() *Queue {
	return &Queue{
		items: make([]string, 16),
		size:  0,
		head:  0,
	}
}
