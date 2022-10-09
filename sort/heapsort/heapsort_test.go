package heapsort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
)

func TestHeapSort(t *testing.T) {
	t.Run("Sort a sorted array", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		result := HeapSort(array)
		for i := 0; i < len(result); i++ {
			if result[i] != i+1 {
				t.Errorf("Expected %d, got %d", i+1, result[i])
			}
		}
	})
	t.Run("Sort an array with negative values", func(t *testing.T) {
		array := []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}
		result := HeapSort(array)
		for i := 0; i < len(result); i++ {
			if result[i] != i-9 {
				t.Errorf("Expected %d, got %d", i-9, result[i])
			}
		}
	})
	t.Run("Sort an empty array", func(t *testing.T) {
		array := []int{}
		result := HeapSort(array)
		if len(result) != 0 {
			t.Errorf("Expected %d, got %d", 0, len(result))
		}
	})
	t.Run("Sort a nil array", func(t *testing.T) {
		var array []int
		result := HeapSort(array)
		if len(result) != 0 {
			t.Errorf("Expected %d, got %d", 0, len(result))
		}
	})
	t.Run("Sort an array with duplicates", func(t *testing.T) {
		array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1}
		result := []int{1, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		output := HeapSort(array)
		for i := 0; i < len(output); i++ {
			if output[i] != result[i] {
				t.Errorf("Expected %d, got %d", result[i], output[i])
			}
		}

	})
	t.Run("Sort an array with random values", func(t *testing.T) {
		array := make([]int, 10)
		for i := 0; i < 10; i++ {
			array[i] = rand.Intn(100)
		}
		output := HeapSort(array)
		sort.Ints(array)
		for i := 0; i < len(output); i++ {
			if output[i] != array[i] {
				t.Errorf("Expected %d, got %d", array[i], output[i])
			}
		}
	})

	t.Run("Test with a large array", func(t *testing.T) {
		array := make([]int, 100)
		for i := 0; i < 100; i++ {
			array[i] = rand.Intn(100)
		}
		result := HeapSort(array)
		sort.Ints(array)
		for i := 0; i < len(result); i++ {
			if result[i] != array[i] {
				t.Errorf("Expected %d, got %d", array[i], result[i])
			}
		}
	})
}

var inputSize = []int{1000000}

func BenchmarkHeapSort(b *testing.B) {
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			array := make([]int, size)
			for i := 0; i < size; i++ {
				array[i] = rand.Intn(size)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				HeapSort(array)
			}
		})
	}
}
