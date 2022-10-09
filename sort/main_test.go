package main

import (
	"fmt"
	"github.com/corentings/golang-exercices/sort/countingsort"
	"github.com/corentings/golang-exercices/sort/mergesort"
	"github.com/corentings/golang-exercices/sort/utils"
	"math/rand"
	"testing"
	"time"
)

var inputSize = []int{10, 100, 1000, 10000, 1000000}

func benchmarkSort(b *testing.B, sort func([]int) []int) {
	seed, err := utils.GenerateRandomSeed()
	if err != nil || seed == -1 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(seed)
	}

	for _, size := range inputSize {
		array := make([]int, size)
		for i := 0; i < size; i++ {
			array[i] = rand.Intn(size)
		}
		b.Run(fmt.Sprintf("input_size_%d", size), func(b *testing.B) {
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sort(array)
			}
		})
	}
}

func Benchmark_CountingSort(b *testing.B) {
	benchmarkSort(b, countingsort.CountingSort)
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
