package linkedList

type Node struct {
	value int
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Size() int {
	size := 0
	for n := l.head; n != nil; n = n.next {
		size++
	}
	return size
}

func (l *LinkedList) Empty() bool {
	return l.head == nil
}

func (l *LinkedList) ValueAt(index int) int {
	node := l.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.value
}

func (l *LinkedList) PushFront(value int) {
	l.head = &Node{value, l.head}
}

func (l *LinkedList) PopFront() {
	l.head = l.head.next
}

func (l *LinkedList) PushBack(value int) {
	if l.Empty() {
		l.PushFront(value)
		return
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	node.next = &Node{value, nil}
}

func (l *LinkedList) PopBack() {
	switch {
	case l.Empty():
		return
	case l.Size() == 1:
		l.head = nil
	default:
		node := l.head
		for node.next.next != nil {
			node = node.next
		}
		node.next = nil
	}
}

func (l *LinkedList) Front() int {
	if l.Empty() {
		return 0
	}
	return l.head.value
}

func (l *LinkedList) Back() int {
	if l.Empty() {
		return 0
	}
	node := l.head
	for node.next != nil {
		node = node.next
	}
	return node.value
}

func (l *LinkedList) Insert(index, value int) {
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
		node.next = &Node{value, node.next}

	default:
		return
	}
}

func (l *LinkedList) Erase(index int) {
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
	default:
		return
	}

}

func (l *LinkedList) ValueNFromEnd(n int) int {
	node := l.head
	for i := 0; i < l.Size()-n-1; i++ {
		node = node.next
	}
	return node.value
}
