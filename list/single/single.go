package single

import (
	"errors"
)

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	next *Node
	size int
}

func New(values ...interface{}) *List {
	list := &List{}
	if values != nil {
		list.PushFront(values...)
	}
	return list
}

func (l *List) PushFront(values ...interface{}) {
	for _, value := range values {
		node := &Node{value, l.next}
		l.next = node
		l.size++
	}
}

func (l *List) Empty() bool {
	return l.size == 0
}

func (l *List) ValueAt(index int) (interface{}, error) {
	if index > l.size {
		return nil, errors.New("Index is out of list range")
	}
	currentNode := l.next
	for i := 0; i != index; i++ {
		currentNode = currentNode.next
	}
	return currentNode.value, nil
}

func (l *List) PopFront() (interface{}, error) {
	if l.Empty() {
		return nil, errors.New("Can't call PopFront when list is empty")
	}
	popped := l.next.value
	l.next = l.next.next
	return popped, nil
}

func (l *List) PushBack(value interface{}) {
	n := &Node{value, nil}
	b := l.Back()
	b.next = n
	l.size++
}

func (l *List) Front() *Node {
	return l.next
}

func (l *List) Back() *Node {
	currentNode := l.next
	for i := 0; i < l.size-1; i++ {
		currentNode = currentNode.next
	}
	return currentNode
}

func (l *List) Display() {
	currentNode := l.next
	for i := 0; i < l.size; i++ {
		currentNode = currentNode.next
	}
}
