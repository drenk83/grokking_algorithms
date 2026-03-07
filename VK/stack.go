package main

import "errors"

var ErrStackEmpty = errors.New("stack is empty")

type Stack struct {
	items []string
}

func (s *Stack) Push(elem string) {
	s.items = append(s.items, elem)
}

func (s *Stack) Pop() (string, error){
	size := len(s.items)
	if size == 0 {
		return "", ErrStackEmpty
	}

	out := s.items[size - 1]
	s.items = s.items[:size - 1]
	return out, nil
}

func (s *Stack) Peek() (string, error){
	size := len(s.items)
	if size == 0 {
		return "", ErrStackEmpty
	}

	return s.items[size - 1], nil
}

func (s *Stack) Size() int {
	return len(s.items)
}

func NewStack() *Stack {
	arr := make([]string, 0, 16)
	return &Stack{
		items: arr,
	}
}
