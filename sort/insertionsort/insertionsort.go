package insertionsort

import "github.com/corentings/golang-exercices/sort/utils"

func Insertionsort[N utils.Numeric](array []N) []N {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && array[j] < array[j-1]; j-- {
			utils.Swap(&array, j, j-1)
		}
	}
	return array
}
