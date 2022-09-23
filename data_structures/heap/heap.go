package heap

type MinHeap []int

func (h *MinHeap) New(array []int) {
	for _, v := range array {
		h.Push(v)
	}
}

func (h *MinHeap) Clear() {
	*h = nil
}

func (h *MinHeap) Display() {
	for i := 0; i < h.Len(); i++ {
		print((*h)[i], " ")
	}
	println()
}

func (h *MinHeap) Leaf(index int) bool {
	return index >= h.Len()/2 && index <= h.Len()-1
}

func (h *MinHeap) GetParent(index int) (int, bool) {
	if index == 0 || index >= h.Len() {
		return 0, false
	}

	return (*h)[(index-1)/2], true
}

func (h *MinHeap) GetLeftChild(index int) (int, bool) {
	left := 2*index + 1
	if left >= h.Len() {
		return 0, false
	}

	return (*h)[left], true
}

func (h *MinHeap) GetRightChild(index int) (int, bool) {
	right := 2*index + 2
	if right >= h.Len() {
		return 0, false
	}

	return (*h)[right], true
}

func (h *MinHeap) Len() int {
	return len(*h)
}

func (h *MinHeap) IsEmpty() bool {
	return h.Len() == 0
}

func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i] < (*h)[j]
}

func (h *MinHeap) Peek() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	return (*h)[0], true
}

func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MinHeap) Push(value int) {
	*h = append(*h, value)
	h.BubbleUp(h.Len() - 1)
}

func (h *MinHeap) Pop() (int, bool) {
	if h.IsEmpty() {
		return 0, false
	}

	v := (*h)[0]
	(*h)[0] = 0 // avoid memory leak
	*h = (*h)[1:]
	h.BubbleDown(0, h.Len())
	return v, true
}

func (h *MinHeap) BubbleUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if !h.Less(index, parentIndex) {
			break
		}
		h.Swap(index, parentIndex)
		index = parentIndex
	}
}

func (h *MinHeap) BubbleDown(index int, size int) {
	for index < size {
		leftIndex := 2*index + 1
		rightIndex := 2*index + 2
		smallestIndex := index
		if (leftIndex < size && h.Less(leftIndex, index)) || ((rightIndex < size) && h.Less(rightIndex, index)) {
			if h.Less(leftIndex, rightIndex) && leftIndex < size {
				h.Swap(index, leftIndex)
				smallestIndex = leftIndex
			} else {
				h.Swap(index, rightIndex)
				smallestIndex = rightIndex
			}
		}
		if smallestIndex == index {
			break
		}
		index = smallestIndex
	}
}
