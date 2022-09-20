package linkedList

type LinkedLister interface {
	Empty() bool
	Size() int
	ValueAt(index int) int
	PushFront(value int)
	PushBack(value int)
	PopFront()
	PopBack()
	Front() int
	Back() int
	Insert(index, value int)
	Erase(index int)
	ValueNFromEnd(n int) int
}
