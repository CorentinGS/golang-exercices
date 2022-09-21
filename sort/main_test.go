package main

import (
	"fmt"
	"github.com/corentings/golang-exercices/sort/mergesort"
	"math/rand"
	"sort"
	"testing"
)

var inputSize = []int{1000000}
var array = make([]int, inputSize[0])

func init() {
	for i := 0; i < inputSize[0]; i++ {
		array[i] = rand.Intn(inputSize[0])
	}
}

func benchmarkSort(b *testing.B, sort func([]int) []int) {
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sort(array)
			}
		})
	}
}

func Benchmark_ParallelMergesort(b *testing.B) {
	benchmarkSort(b, mergesort.ParallelMergesort)
}

func Benchmark_Mergesort(b *testing.B) {
	benchmarkSort(b, mergesort.Mergesort)
}

func Benchmark_MergeSortInsertion(b *testing.B) {
	benchmarkSort(b, mergesort.MergeSortInsertion)
}

/*
func Benchmark_BubbleSort(b *testing.B) {
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				bubblesort.Bubblesort(array)
			}
		})
	}
}


func Benchmark_InsertionSort(b *testing.B) {
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				insertionsort.Insertionsort(array)
			}
		})
	}
}
*/

func Benchmark_NativeSort(b *testing.B) {
	for _, size := range inputSize {
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sort.Ints(array)
			}
		})
	}
}
