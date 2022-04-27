package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
	prev  *ListNode
}

type DLinkedList struct {
	sentinel *ListNode
	size     int
}

func (n *ListNode) InsertAfter(v int) {
	n.next = &ListNode{v, n.next, n.prev}
}

func (l *DLinkedList) InsertFront(v int) {
	l.sentinel.prev = &ListNode{v, l.sentinel, l.sentinel.prev}
}

func NewLinkedList(v int) *DLinkedList {
	l := &DLinkedList{&ListNode{0, nil, nil}, 0}
	node := &ListNode{v, l.sentinel, l.sentinel}
	l.sentinel.next = node
	l.sentinel.prev = node
	return l
}

func main() {
	list := NewLinkedList(2)
	list.InsertFront(3)
	list.InsertFront(1)
	fmt.Println(list.sentinel.prev)
}
