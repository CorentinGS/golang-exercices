package countingsort

func CountingSort(arr []int) []int {
	max := 0
	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	count := make([]int, max+1)
	for _, v := range arr {
		count[v]++
	}

	z := 0
	for i, c := range count {
		for ; c > 0; c-- {
			arr[z] = i
			z++
		}
	}

	return arr
}
