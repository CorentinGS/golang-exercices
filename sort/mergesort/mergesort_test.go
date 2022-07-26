package mergesort

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func generateRandomArray(size int) []int {
	input := make([]int, size)
	for i := 0; i < size; i++ {
		input[i] = rand.Intn(size)
	}
	return input
}

func Test_merge(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "Test merge",
			args: args{
				left:  []int{1, 2, 3, 4, 5, 6, 7, 8},
				right: []int{},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "Test merge 2",
			args: args{
				right: []int{7, 6, 5, 4, 3, 2, 1},
				left:  []int{8},
			},
			expected: []int{7, 6, 5, 4, 3, 2, 1, 8},
		},
		{
			name: "Test merge 3",
			args: args{
				left:  []int{3, 4, 5, 6},
				right: []int{2, 4, 6, 8, 10},
			},
			expected: []int{2, 3, 4, 4, 5, 6, 6, 8, 10},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := merge(tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("merge() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_mergesort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name     string
		args     args
		expected []int
	}{
		{
			name: "Test Mergesort",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "Test Mergesort 2",
			args: args{
				array: []int{8, 7, 6, 5, 4, 3, 2, 1},
			},
			expected: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "Test Mergesort 3",
			args: args{
				array: []int{69, 5, 6, 2, 64, 4687, 6, 54, 564, 1, 45, 4, 54, 2, 4, 41, 54, 5, 4268, 4, 514, 564, 4, 2564, 814, 6, 41, 6514, 68, 0, 5, 42, 67, 28642},
			},
			expected: []int{0, 1, 2, 2, 4, 4, 4, 4, 5, 5, 5, 6, 6, 6, 41, 41, 42, 45, 54, 54, 54, 64, 67, 68, 69, 514, 564, 564, 814, 2564, 4268, 4687, 6514, 28642},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mergesort(tt.args.array); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Mergesort() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func Test_mergesort_random(t *testing.T) {
	input := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		input[i] = rand.Intn(1000000)
	}

	got := Mergesort(input)
	for i := 0; i < len(got)-1; i++ {
		if got[i] > got[i+1] {
			t.Errorf("Mergesort() = %v, want %v", got, input)
		}
	}
}

func Test_mergesortinsertion(t *testing.T) {
	numberOfTests := 100

	t.Run("Test MergeSortInsertion", func(t *testing.T) {
		for i := 0; i < numberOfTests; i++ {
			array := generateRandomArray(numberOfTests)
			expected := make([]int, len(array))
			copy(expected, array)
			sort.Ints(expected)
			if got := MergeSortInsertion(array); !reflect.DeepEqual(got, expected) {
				t.Errorf("MergeSortInsertion() = %v, want %v", got, expected)
			}
		}
	})

	t.Run("Test MergeSortInsertion with empty array", func(t *testing.T) {
		var array []int
		var expected []int
		if got := MergeSortInsertion(array); !reflect.DeepEqual(got, expected) {
			t.Errorf("MergeSortInsertion() = %v, want %v", got, expected)
		}
	})

	t.Run("Test MergeSortInsertion with one element", func(t *testing.T) {
		array := []int{1}
		expected := []int{1}
		if got := MergeSortInsertion(array); !reflect.DeepEqual(got, expected) {
			t.Errorf("Bubblesort() = %v, want %v", got, expected)
		}
	})

}

func Test_parallelmergesort_random(t *testing.T) {
	input := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		input[i] = rand.Intn(1000000)
	}

	got := ParallelMergesort(input)
	for i := 0; i < len(got)-1; i++ {
		if got[i] > got[i+1] {
			t.Errorf("Mergesort() = %v, want %v", got, input)
		}
	}
}

func Benchmark_mergesort(b *testing.B) {
	inputSize := []int{1000000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			input := make([]int, size)
			for i := 0; i < size; i++ {
				input[i] = rand.Intn(size)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				Mergesort(input)
			}
		})
	}
}

func Benchmark_parallelmergesort(b *testing.B) {
	inputSize := []int{1000000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			input := make([]int, size)
			for i := 0; i < size; i++ {
				input[i] = rand.Intn(size)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				ParallelMergesort(input)
			}
		})
	}
}

func Benchmark_mergesortinsertion(b *testing.B) {
	inputSize := []int{1000000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			input := make([]int, size)
			for i := 0; i < size; i++ {
				input[i] = rand.Intn(size)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				MergeSortInsertion(input)
			}
		})
	}
}

func Benchmark_native_sort(b *testing.B) {
	inputSize := []int{1000000}
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			input := make([]int, size)
			for i := 0; i < size; i++ {
				input[i] = rand.Intn(size)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sort.Ints(input)
			}
		})
	}
}

func Fuzz_mergesort(f *testing.F) {
	testCases := []int{51, 634, 9, 8941, 354, 0}
	slice := []int{69, 5, 6, 2, 64, 4687, 6, 54, 564, 1, 45, 4, 54, 2, 4, 41, 54, 5, 4268, 4, 514, 564, 4, 2564, 814, 6, 41, 6514, 68, 0, 5, 42, 67, 28642}

	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, i int) {
		slice := append(slice, i)
		got := Mergesort(slice)
		for i := 0; i < len(got)-1; i++ {
			if got[i] > got[i+1] {
				t.Errorf("Mergesort() = %v, want %v", got, slice)
			}
		}
	})
}

func Fuzz_mergesortparallel(f *testing.F) {
	testCases := []int{51, 634, 9, 8941, 354, 0}
	slice := []int{69, 5, 6, 2, 64, 4687, 6, 54, 564, 1, 45, 4, 54, 2, 4, 41, 54, 5, 4268, 4, 514, 564, 4, 2564, 814, 6, 41, 6514, 68, 0, 5, 42, 67, 28642}

	for _, tc := range testCases {
		f.Add(tc)
	}
	f.Fuzz(func(t *testing.T, i int) {
		slice := append(slice, i)
		got := ParallelMergesort(slice)
		for i := 0; i < len(got)-1; i++ {
			if got[i] > got[i+1] {
				t.Errorf("ParallelMergesort() = %v, want %v", got, slice)
			}
		}
	})
}
