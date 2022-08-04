package double

import (
	"errors"
)

type Node[T any] struct {
	value T
	Next  *Node[T]
	Prev  *Node[T]
}

type List[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

func New(values ...any) *List[any] {
	list := List[any]{}
	if len(values) != 0 {
		fn := &Node[any]{values[0], nil, nil}
		list.Head = fn
		list.Tail = fn
		list.size++
		for _, value := range values[1:] {
			node := &Node[any]{value, nil, list.Tail}
			list.Tail = node
			list.size++
		}
	}
	return &list
}

func (l *List[any]) PushFront(values ...any) {
	for _, value := range values {
		node := &Node[any]{value, nil, l.Tail}
		l.Tail = node
		if l.Head == nil {
			l.Head = node
		}
		l.size++
	}
}

func (l *List[T]) Empty() bool {
	return l.size == 0
}

func (l *List[T]) Size() int {
	return l.size
}

func (l *List[T]) ValueAt(index int) (*Node[T], error) {
	if index > l.size || index < 1 {
		return nil, errors.New("Index is out of list range")
	}

	currentNode := l.Head
	for i := 1; i < index; i++ {
		currentNode = currentNode.Next
	}
	return currentNode, nil
}

func (l *List[T]) PopFront() (interface{}, error) {
	if l.Empty() {
		return nil, errors.New("Can't call PopFront when list is empty")
	}
	p := l.Head.value
	l.Head = l.Head.Next
	return p, nil
}

func (l *List[T]) PushBack(value T) {
	n := &Node[T]{value, l.Tail, l.Tail.Prev}
	b := l.Tail
	b.Next = n
	l.size++
}
