package heap

import "fmt"

type MaxHeap []int

func New(arr []int) MaxHeap {
	h := make(MaxHeap, len(arr))
	h.BuildHeap(arr)
	return h
}

func (h *MaxHeap) ToArray() []int {
	return *h
}

func (h *MaxHeap) BuildHeap(arr []int) {
	copy(*h, arr)
	for i := h.Len() / 2; i >= 0; i-- {
		h.Heapify(h.Len(), i)
	}
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Peek() (int, error) {
	if h.IsEmpty() {
		return 0, fmt.Errorf("heap is empty")
	}
	item := (*h)[0]
	(*h)[0] = (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	h.Heapify(h.Len(), 0)
	return item, nil
}

func (h *MaxHeap) Len() int {
	return len(*h)
}

func (h *MaxHeap) HeapifyUp(index int) {
	for h.GetParentIndex(index) >= 0 && (*h)[h.GetParentIndex(index)] < (*h)[index] {
		(*h)[h.GetParentIndex(index)], (*h)[index] = (*h)[index], (*h)[h.GetParentIndex(index)]
		index = h.GetParentIndex(index)
	}
}

func (h *MaxHeap) Push(value int) {
	*h = append(*h, value)
	h.HeapifyUp(h.Len() - 1)
}

func (h *MaxHeap) IsEmpty() bool {
	return h.Len() == 0
}

func (h *MaxHeap) ExtractPeek() (int, error) {
	if h.IsEmpty() {
		return 0, fmt.Errorf("heap is empty")
	}
	return (*h)[0], nil
}

func (h *MaxHeap) Heapify(size int, i int) {
	largest := i
	left := h.GetLeftChildIndex(i)
	right := h.GetRightChildIndex(i)

	if left < size && (*h)[left] > (*h)[largest] {
		largest = left
	}

	if right < size && (*h)[right] > (*h)[largest] {
		largest = right
	}

	if largest != i {
		(*h)[i], (*h)[largest] = (*h)[largest], (*h)[i]
		h.Heapify(size, largest)
	}

}

func (h *MaxHeap) GetParentIndex(i int) int {
	return (i - 1) / 2
}

func (h *MaxHeap) GetLeftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *MaxHeap) GetRightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}
