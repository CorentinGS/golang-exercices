package heapsort

import "github.com/corentings/golang-exercices/data_structures/heap"

func HeapSort(arr []int) []int {
	h := heap.New()

	h.BuildHeap(arr)

	for i := h.Len() - 1; i >= 0; i-- {
		h.Swap(0, i)
		h.Heapify(i, 0)
	}

	return h.ToArray()
}
