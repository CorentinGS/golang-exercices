package utils

func Swap[N Numeric](array *[]N, i int, j int) {
	(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
}
