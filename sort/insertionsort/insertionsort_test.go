package insertionsort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func generateRandomArray(size int) []int {
	array := make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Intn(size)
	}
	return array
}

func generateRandomArrayFloat(size int) []float64 {
	array := make([]float64, size)
	for i := 0; i < size; i++ {
		array[i] = rand.Float64()
	}
	return array
}

func generateRandomArrayUint(size int) []uint {
	array := make([]uint, size)
	for i := 0; i < size; i++ {
		array[i] = uint(rand.Uint32())
	}
	return array
}

func BenchmarkInsertionSort(b *testing.B) {
	inputSize := []int{0, 10, 100}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			input := generateRandomArray(size)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Insertionsort(input)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	numberOfTests := 100

	t.Run("Test Insertionsort", func(t *testing.T) {
		for i := 0; i < numberOfTests; i++ {
			array := generateRandomArray(numberOfTests)
			expected := make([]int, len(array))
			copy(expected, array)
			sort.Ints(expected)
			if got := Insertionsort(array); !reflect.DeepEqual(got, expected) {
				t.Errorf("Insertionsort() = %v, want %v", got, expected)
			}
		}
	})

	t.Run("Test Insertionsort with empty array", func(t *testing.T) {
		var array []int
		var expected []int
		if got := Insertionsort(array); !reflect.DeepEqual(got, expected) {
			t.Errorf("Insertionsort() = %v, want %v", got, expected)
		}
	})

	t.Run("Test Insertionsort with one element", func(t *testing.T) {
		array := []int{1}
		expected := []int{1}
		if got := Insertionsort(array); !reflect.DeepEqual(got, expected) {
			t.Errorf("Insertionsort() = %v, want %v", got, expected)
		}
	})

	t.Run("Test Insertionsort with float array", func(t *testing.T) {
		for i := 0; i < numberOfTests; i++ {
			array := generateRandomArrayFloat(numberOfTests)
			expected := make([]float64, len(array))
			copy(expected, array)
			sort.Float64s(expected)
			if got := Insertionsort(array); !reflect.DeepEqual(got, expected) {
				t.Errorf("Insertionsort() = %v, want %v", got, expected)
			}
		}
	})

	t.Run("Test Insertionsort with uint array", func(t *testing.T) {
		for i := 0; i < numberOfTests; i++ {
			array := generateRandomArrayUint(numberOfTests)
			expected := make([]uint, len(array))
			copy(expected, array)
			sort.Slice(expected, func(i, j int) bool { return expected[i] < expected[j] })
			if got := Insertionsort(array); !reflect.DeepEqual(got, expected) {
				t.Errorf("Insertionsort() = %v, want %v", got, expected)
			}
		}
	})
}
