package countingsort

import "testing"

func Test_CountingSort(t *testing.T) {
	arr := []int{1, 4, 1, 2, 7, 5, 2}
	arr = CountingSort(arr)
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			t.Errorf("CountingSort failed. Got %v, expected %v", arr, []int{1, 1, 2, 2, 4, 5, 7})
			break
		}
	}
}
