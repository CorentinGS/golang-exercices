package heapsort

import "github.com/corentings/golang-exercices/data_structures/heap"

func HeapSort(array []int) {
	heap := heap.MinHeap{}
	heap.New(array)
	for i := heap.Len() - 1; i > 0; i-- {
		heap.Swap(0, i)
		heap.BubbleDown(0, i)
	}
}
