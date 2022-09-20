package bubblesort

import "github.com/corentings/golang-exercices/sort/utils"

func Bubblesort[N utils.Numeric](array []N) []N {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				utils.Swap(&array, j, j+1)
			}
		}
	}
	return array
}
