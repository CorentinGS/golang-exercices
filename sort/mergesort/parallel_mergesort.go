package mergesort

func parallelMergesort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	if len(slice) < 2048 {
		return mergesort(slice)
	} else {
		done := make(chan struct{})
		defer close(done)

		middle := len(slice) / 2
		var left, right []int
		go func() {
			left = parallelMergesort(slice[:middle])
			done <- struct{}{}
		}()
		right = parallelMergesort(slice[middle:])
		<-done
		return merge(left, right)
	}

	return slice
}

func parallelMergeSortTry(slice []int, max int) []int {
	if len(slice) < 2 {
		return slice
	}
	if len(slice) < max {
		return mergesort(slice)
	} else {
		done := make(chan struct{})
		defer close(done)

		middle := len(slice) / 2
		var left, right []int
		go func() {
			left = parallelMergeSortTry(slice[:middle], max)
			done <- struct{}{}
		}()
		right = parallelMergeSortTry(slice[middle:], max)
		<-done
		return merge(left, right)
	}

	return slice
}
