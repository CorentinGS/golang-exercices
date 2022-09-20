package linkedList

type DoubleNode struct {
	value int
	next  *DoubleNode
	prev  *DoubleNode
}

func (n *DoubleNode) Next() *DoubleNode {
	return n.next
}

func (n *DoubleNode) Prev() *DoubleNode {
	return n.prev
}

func (n *DoubleNode) Value() int {
	return n.value
}

type DoubleLinkedList struct {
	head *DoubleNode
	tail *DoubleNode
}

func (l *DoubleLinkedList) Empty() bool {
	return l.head == nil
}

func (l *DoubleLinkedList) Size() int {
	size := 0
	for n := l.head; n != nil; n = n.next {
		size++
	}
	return size
}

func (l *DoubleLinkedList) ValueAt(index int) int {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value
}

func (l *DoubleLinkedList) PushFront(value int) {
	l.head = &DoubleNode{value, l.head, nil}
	if l.tail == nil {
		l.tail = l.head
	}
}

func (l *DoubleLinkedList) PopFront() {
	switch {
	case l.Empty():
		return
	case l.Size() == 1:
		l.head = nil
		l.tail = nil
	default:
		l.head = l.head.next
		l.head.prev = nil
	}

}

func (l *DoubleLinkedList) PushBack(value int) {
	if l.Empty() {
		l.PushFront(value)
		return
	}
	l.tail.next = &DoubleNode{value, nil, l.tail}
	l.tail = l.tail.next
}

func (l *DoubleLinkedList) PopBack() {
	switch {
	case l.Empty():
		return
	case l.Size() == 1:
		l.head = nil
		l.tail = nil
	default:
		l.tail = l.tail.prev
		l.tail.next = nil
	}
}

func (l *DoubleLinkedList) Front() int {
	if l.Empty() {
		return 0
	}
	return l.head.value
}

func (l *DoubleLinkedList) Back() int {
	if l.Empty() {
		return 0
	}
	return l.tail.value
}

func (l *DoubleLinkedList) Insert(index, value int) {
	switch {
	case l.Empty():
		l.PushFront(value)
	case index == 0:
		l.PushFront(value)
	case index == l.Size() || index == -1:
		l.PushBack(value)
	case index > 0 && index < l.Size():
		node := l.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		node.next = &DoubleNode{value, node.next, node}

	default:
		return
	}
}

func (l *DoubleLinkedList) Erase(index int) {
	switch {
	case l.Empty():
		return
	case index == 0:
		l.PopFront()
	case index == l.Size()-1 || index == -1:
		l.PopBack()
	case index < -1 && index > (l.Size()-1)*-1:
		l.Erase(l.Size() + index)
	case index < l.Size()-1 && index > 0:
		node := l.head
		for i := 0; i < index-1; i++ {
			node = node.next
		}
		node.next = node.next.next
		node.next.prev = node

	default:
		return
	}

}

func (l *DoubleLinkedList) ValueNFromEnd(n int) int {
	node := l.head
	for i := 0; i < l.Size()-n-1; i++ {
		node = node.next
	}
	return node.value
}
