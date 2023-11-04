package mysinglylinkedlist

type MySinglyLinkedList struct {
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func NewMySinglyLinkedList() *MySinglyLinkedList {
	return &MySinglyLinkedList{}
}

func (l *MySinglyLinkedList) Len() int {
	length := 0
	for n := l.head; n != nil; n = n.next {
		length++
	}

	return length
}

func (l *MySinglyLinkedList) Get(index int) int {
	n := l.head
	if index < 0 || n == nil {
		panic("index out of range")
	}

	for i := 0; i < index; i++ {
		n = n.next
		if n == nil {
			panic("index out of range")
		}
	}

	return n.value
}

func (l *MySinglyLinkedList) GetAll() []int {
	result := []int{}

	for n := l.head; n != nil; n = n.next {
		result = append(result, n.value)
	}

	return result
}

func (l *MySinglyLinkedList) AddToHead(value int) {
	l.head = &Node{value: value, next: l.head}
}

func (l *MySinglyLinkedList) AddToTail(value int) {
	if l.head == nil {
		l.head = &Node{value: value}
		return
	}

	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{value: value}
}

func (l *MySinglyLinkedList) AddAtIndex(index, value int) {
	if index < 0 {
		panic("index out of range")
	}

	if index == 0 {
		l.AddToHead(value)
		return
	}

	node := l.head
	for i := 0; i < index-1; i++ {
		if node == nil {
			panic("index out of range")
		}
		node = node.next
	}
	node.next = &Node{value: value, next: node.next}
}

func (l *MySinglyLinkedList) Remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	node := l.head
	for node.next != nil {
		if node.next.value == value {
			node.next = node.next.next
			return
		}
		node = node.next
	}
}

func (l *MySinglyLinkedList) RemoveFromHead() {
	if l.head == nil {
		return
	}

	l.head = l.head.next
}

func (l *MySinglyLinkedList) RemoveFromTail() {
	if l.head == nil {
		return
	}

	if l.head.next == nil {
		l.head = nil
		return
	}

	node := l.head
	for node.next.next != nil {
		node = node.next
	}
	node.next = nil
}

func (l *MySinglyLinkedList) RemoveAtIndex(index int) {
	if index < 0 {
		panic("index out of range")
	}

	if index == 0 {
		l.RemoveFromHead()
		return
	}

	node := l.head
	for i := 0; i < index-1; i++ {
		if node == nil {
			panic("index out of range")
		}
		node = node.next
	}
	node.next = node.next.next
}