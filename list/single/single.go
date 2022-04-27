package single

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
