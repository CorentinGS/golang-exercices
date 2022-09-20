package mergesort

import "github.com/corentings/golang-exercices/sort/insertionsort"

const K = 10

func Mergesort(array []int) []int {
	size := len(array)
	if size < 2 {
		return array
	}

	middle := size / 2

	return merge(Mergesort(array[:middle]), Mergesort(array[middle:]))
}

func MergeSortInsertion(array []int) []int {
	size := len(array)
	if size < 2 {
		return array
	}

	if size < K {
		return insertionsort.Insertionsort(array)
	}

	middle := size / 2

	return merge(MergeSortInsertion(array[:middle]), MergeSortInsertion(array[middle:]))
}

func merge(left, right []int) []int {

	slice := make([]int, len(left)+len(right))

	i, j := 0, 0

	for k := 0; k < cap(slice); k++ {
		switch {
		case i >= len(left):
			slice[k] = right[j]
			j++

		case j >= len(right):
			slice[k] = left[i]
			i++

		case left[i] < right[j]:
			slice[k] = left[i]
			i++

		default:
			slice[k] = right[j]
			j++
		}
	}

	return slice
}
