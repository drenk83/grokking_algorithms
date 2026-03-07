package main

type Node struct {
	data string
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) MakeNewList() *LinkedList {
	return &LinkedList{
		head: nil,
		size: 0,
	}
}

func (l *LinkedList) AddNewHead(elem string) {
	newNode := &Node{
		data: elem,
		next: l.head,
	}
	l.head = newNode
	l.size++
}

func (l *LinkedList) AddNewTail(elem string) {
	newNode := &Node{
		data: elem,
		next: nil,
	}

	if l.size == 0 {
		l.head = newNode
		l.size++
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	l.size++
}

func (l *LinkedList) Insert(after, elem string) {
	current := l.head

	// Здесь порядок очень важен, так как при currnet == nil
	// он не будет проверь второе условие (а current.data при nil вызовет панику)
	for current != nil && current.data != after {
		current = current.next
	}

	if current == nil {
		return
	}

	newNode := &Node{
		data: elem,
		next: current.next,
	}
	current.next = newNode
	l.size++
}

func (l *LinkedList) Search(elem string) *Node {
	current := l.head

	for current != nil {
		if current.data == elem {
			return current
		}
		current = current.next
	}
	return nil
}

func (l *LinkedList) Delete(elem string) {
	current := l.head

	if current != nil && current.data == elem {
		l.head = current.next
		l.size--
		return
	}

	for current != nil && current.next != nil {
		if current.next.data == elem {
			current.next = current.next.next
			l.size--
			return
		}
		current = current.next
	}
}

func main() {

}
