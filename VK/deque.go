package main

import "errors"

var ErrEmpty = errors.New("deque is empty")

type Node struct {
	data string
	prev *Node
	next *Node
}

type Deque struct {
	size int
	head *Node
	tail *Node
}

func (d *Deque) PushFront(elem string) {
	newNode := &Node{
		data: elem,
		prev: nil,
		next: d.head,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
		d.size++
		return
	}

	d.head.prev = newNode
	d.head = newNode
	d.size++
}

func (d *Deque) PushBack(elem string) {
	newNode := &Node{
		data: elem,
		prev: d.tail,
		next: nil,
	}

	if d.tail == nil {
		d.head = newNode
		d.tail = newNode
		d.size++
		return
	}

	d.tail.next = newNode
	d.tail = newNode
	d.size++
}

func (d *Deque) PopFront() (string, error) {
	if d.size == 0 {
		return "", ErrEmpty
	}

	if d.size == 1 {
		out := d.head.data
		d.head = nil
		d.tail = nil
		d.size--
		return out, nil
	}

	out := d.head.data
	d.head = d.head.next
	d.head.prev = nil
	d.size--
	return out, nil
}

func (d *Deque) PopBack() (string, error) {
	if d.size == 0 {
		return "", ErrEmpty
	}

	if d.size == 1 {
		out := d.head.data
		d.head = nil
		d.tail = nil
		d.size--
		return out, nil
	}

	out := d.tail.data
	d.tail = d.tail.prev
	d.tail.next = nil
	d.size--
	return out, nil
}

func NewDeque() *Deque {
	return &Deque{
		size: 0,
		head: nil,
		tail: nil,
	}
}
