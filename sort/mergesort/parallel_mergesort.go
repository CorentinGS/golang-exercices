package mergesort

func ParallelMergesort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	if len(slice) < 2048 {
		return MergeSortInsertion(slice)
	} else {
		done := make(chan struct{})
		defer close(done)

		middle := len(slice) / 2
		var left, right []int
		go func() {
			left = ParallelMergesort(slice[:middle])
			done <- struct{}{}
		}()
		right = ParallelMergesort(slice[middle:])
		<-done
		return merge(left, right)
	}
}
