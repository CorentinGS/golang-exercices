package heap

func HeapSort(array []int) MinHeap {
	heap := MinHeap{}
	heap.New(array)

	result := make([]int, heap.Len())

	for min, ok := heap.Peek(); ok; min, ok = heap.Peek() {
		result = append(result, min)
		heap.Pop()
	}

	heap.Display()

	return heap
}
