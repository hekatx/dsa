package single

import "errors"

var m = map[string]string{
	"valueAt":  "Index is out of list range",
	"popFront": "Can't call PopFront when list is empty",
	"reverse":  "Index is out of list range",
}

type Node[T any] struct {
	value T
	Next  *Node[T]
}

type Head[T any] struct {
	Next *Node[T]
	size int
}

func New[T any](values []T) Head[T] {
	list := Head[T]{}
	if values != nil {
		list.PushFront(values)
	}
	return list
}

func (h *Head[T]) PushFront(values []T) {
	for _, value := range values {
		node := &Node[T]{value, h.Next}
		h.Next = node
		h.size++
	}
}

func (h *Head[T]) Empty() bool {
	return h.size == 0
}

func (h *Head[T]) Size() int {
	return h.size
}

func (h *Head[T]) ValueAt(index int) (interface{}, error) {
	if index > h.size {
		return nil, errors.New(m["valueAt"])
	}
	currentNode := h.Next
	for i := 0; i != index; i++ {
		currentNode = currentNode.Next
	}
	return currentNode.value, nil
}

func (h *Head[T]) PopFront() (interface{}, error) {
	if h.Empty() {
		return nil, errors.New(m["popFront"])
	}
	popped := h.Next.value
	h.Next = h.Next.Next
	return popped, nil
}

func (h *Head[T]) PushBack(value T) {
	n := &Node[T]{value, nil}
	b := h.Back()
	b.Next = n
	h.size++
}

func (h *Head[T]) Front() *Node[T] {
	return h.Next
}

func (h *Head[T]) Back() *Node[T] {
	currentNode := h.Next
	for i := 0; i < h.size-1; i++ {
		currentNode = currentNode.Next
	}
	return currentNode
}

func (h *Head[T]) Reverse(n *Node[T]) *Node[T] {
	if n.Next == nil {
		h.Next = n
		return n
	}

	cn := h.Reverse(n.Next)
	n.Next = nil
	cn.Next = n

	return n
}
