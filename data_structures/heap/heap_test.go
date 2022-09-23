package heap

import "testing"

func TestMinHeap_New(t *testing.T) {
	t.Run("Ceate a new heap from ordered array", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		if h.Len() != len(array) {
			t.Errorf("Expected %d, got %d", len(array), h.Len())
		}
		for i := 0; i < h.Len(); i++ {
			if (*h)[i] != array[i] {
				t.Errorf("Expected %d, got %d", array[i], (*h)[i])
			}
		}
	})
	t.Run("Ceate a new heap from unordered array", func(t *testing.T) {
		array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
		h := &MinHeap{}
		h.New(array)
		if h.Len() != len(array) {
			t.Errorf("Expected %d, got %d", len(array), h.Len())
		}
	})
	t.Run("Ceate a new heap from empty array", func(t *testing.T) {
		array := []int{}
		h := &MinHeap{}
		h.New(array)
		if h.Len() != len(array) {
			t.Errorf("Expected %d, got %d", len(array), h.Len())
		}
	})
	t.Run("Ceate a new heap from nil array", func(t *testing.T) {
		var array []int
		h := &MinHeap{}
		h.New(array)
		if h.Len() != len(array) {
			t.Errorf("Expected %d, got %d", len(array), h.Len())
		}
	})
	t.Run("Ceate a new heap from array with negative values", func(t *testing.T) {
		array := []int{-9, -8, -7, -6, -5, -4, -3, -2, -1}
		h := &MinHeap{}
		h.New(array)
		if h.Len() != len(array) {
			t.Errorf("Expected %d, got %d", len(array), h.Len())
		}
		for i := 0; i < h.Len(); i++ {
			if (*h)[i] != i-9 {
				t.Errorf("Expected %d, got %d", array[i], (*h)[i])
			}
		}
	})
}

func TestMinHeap_Clear(t *testing.T) {
	t.Run("Clear a heap", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		h.Clear()
		if h.Len() != 0 {
			t.Errorf("Expected %d, got %d", 0, h.Len())
		}
	})
}

func TestMinHeap_IsEmpty(t *testing.T) {
	t.Run("Check if a heap is empty", func(t *testing.T) {
		array := []int{}
		h := &MinHeap{}
		h.New(array)
		if !h.IsEmpty() {
			t.Errorf("Expected %t, got %t", true, h.IsEmpty())
		}
	})

	t.Run("Check if a heap is not empty", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		if h.IsEmpty() {
			t.Errorf("Expected %t, got %t", false, h.IsEmpty())
		}
	})
}

func TestMinHeap_Push(t *testing.T) {
	t.Run("Push a value into a heap", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		h.Push(10)
		if h.Len() != len(array)+1 {
			t.Errorf("Expected %d, got %d", len(array)+1, h.Len())
		}
	})
	t.Run("Push a negative value into a heap", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		h.Push(-10)
		if h.Len() != len(array)+1 {
			t.Errorf("Expected %d, got %d", len(array)+1, h.Len())
		}
		min, _ := h.Peek()
		if min != -10 {
			t.Errorf("Expected %d, got %d", -10, min)
		}
	})
	t.Run("Push a value into an empty heap", func(t *testing.T) {
		array := []int{}
		h := &MinHeap{}
		h.New(array)
		h.Push(10)
		if h.Len() != len(array)+1 {
			t.Errorf("Expected %d, got %d", len(array)+1, h.Len())
		}
		min, _ := h.Peek()
		if min != 10 {
			t.Errorf("Expected %d, got %d", 10, min)
		}
	})
}

func TestMinHeap_Pop(t *testing.T) {
	t.Run("Pop a value from a heap", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		min, _ := h.Pop()
		if min != 1 {
			t.Errorf("Expected %d, got %d", 1, min)
		}
		if h.Len() != len(array)-1 {
			t.Errorf("Expected %d, got %d", len(array)-1, h.Len())
		}

		new_min, _ := h.Peek()
		if new_min != 2 {
			t.Errorf("Expected %d, got %d", 2, new_min)
		}
	})
	t.Run("Pop a value from an empty heap", func(t *testing.T) {
		array := []int{}
		h := &MinHeap{}
		h.New(array)
		min, ok := h.Pop()
		if ok {
			t.Errorf("Expected %s, got %v", "heap is empty", !ok)
		}
		if min != 0 {
			t.Errorf("Expected %d, got %d", 0, min)
		}
	})
	t.Run("Pop a value from a heap with one element", func(t *testing.T) {
		array := []int{1}
		h := &MinHeap{}
		h.New(array)
		min, _ := h.Pop()
		if min != 1 {
			t.Errorf("Expected %d, got %d", 1, min)
		}
		if h.Len() != len(array)-1 {
			t.Errorf("Expected %d, got %d", len(array)-1, h.Len())
		}
	})
	t.Run("Pop a value from an unsorted heap", func(t *testing.T) {
		array := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
		h := &MinHeap{}
		h.New(array)
		h.Push(0)
		min, _ := h.Pop()
		if min != 0 {
			t.Errorf("Expected %d, got %d", 0, min)
		}
		if h.Len() != len(array) {
			t.Errorf("Expected %d, got %d", len(array), h.Len())
		}
	})
}

func TestMinHeap_GetLeftChild(t *testing.T) {
	t.Run("Get left child of a node", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		left_child, _ := h.GetLeftChild(0)
		if left_child != 2 {
			t.Errorf("Expected %d, got %d", 2, left_child)
		}
	})
	t.Run("Get left child of a node that does not exist", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		left_child, ok := h.GetLeftChild(9)
		if ok {
			t.Errorf("Expected %s, got %v", "node does not exist", !ok)
		}
		if left_child != 0 {
			t.Errorf("Expected %d, got %d", 0, left_child)
		}
	})
}

func TestMinHeap_GetRightChild(t *testing.T) {
	t.Run("Get right child of a node", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		right_child, _ := h.GetRightChild(0)
		if right_child != 3 {
			t.Errorf("Expected %d, got %d", 3, right_child)
		}
	})
	t.Run("Get right child of a node that does not exist", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		right_child, ok := h.GetRightChild(9)
		if ok {
			t.Errorf("Expected %s, got %v", "node does not exist", !ok)
		}
		if right_child != 0 {
			t.Errorf("Expected %d, got %d", 0, right_child)
		}
	})
}

func TestMinHeap_GetParent(t *testing.T) {
	t.Run("Get parent of a node", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		parent, _ := h.GetParent(1)
		if parent != 1 {
			t.Errorf("Expected %d, got %d", 0, parent)
		}
	})
	t.Run("Get parent of a node", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		parent, _ := h.GetParent(4)
		if parent != 2 {
			t.Errorf("Expected %d, got %d", 2, parent)
		}
	})
	t.Run("Get parent of a node that does not exist", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		parent, ok := h.GetParent(9)
		if ok {
			t.Errorf("Expected %s, got %v", "node does not exist", !ok)
		}
		if parent != 0 {
			t.Errorf("Expected %d, got %d", 0, parent)
		}
	})
	t.Run("Get parent of the first node", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		parent, ok := h.GetParent(0)
		if ok {
			t.Errorf("Expected %s, got %v", "node does not exist", !ok)
		}
		if parent != 0 {
			t.Errorf("Expected %d, got %d", 0, parent)
		}
	})
}

func TestMinHeap_Leaf(t *testing.T) {
	t.Run("Check if a node is a leaf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		if !h.Leaf(8) {
			t.Errorf("Expected %v, got %v", true, h.Leaf(8))
		}
	})
	t.Run("Check if a node is a leaf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		if h.Leaf(0) {
			t.Errorf("Expected %v, got %v", false, h.Leaf(0))
		}
	})
	t.Run("Check if a node is a leaf", func(t *testing.T) {
		array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		h := &MinHeap{}
		h.New(array)
		if !h.Leaf(7) {
			t.Errorf("Expected %v, got %v", false, h.Leaf(7))
		}
	})
}
