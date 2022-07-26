package linkedList

import (
	"math/rand"
	"testing"
)

func GenerateRandomList(numberOfTests int) LinkedList {
	l := LinkedList{}
	for i := 0; i < numberOfTests; i++ {
		i := rand.Intn(numberOfTests)
		l.PushBack(i)
	}
	return l
}

// TestLinkedList_Back tests the Back method of the LinkedList
func TestLinkedList_Back(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	// Test empty list
	if l.Back() != 0 {
		t.Errorf("Back() = %v, want %v", l.Back(), 0)
	}
	l = GenerateRandomList(numberOfTests)

	for i := 0; i < numberOfTests; i++ {
		l.PushBack(i)
		if l.Back() != i {
			t.Errorf("Back() = %v, want %v", l.Back(), i)
		}
	}
}

// TestLinkedList_Empty tests the Empty method of the LinkedList
func TestLinkedList_Empty(t *testing.T) {
	type fields struct {
		head *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "Empty list",
			fields: fields{
				head: nil,
			},
			want: true,
		},
		{
			name: "Non-empty list",
			fields: fields{
				head: &Node{
					value: 1,
					next:  nil,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head: tt.fields.head,
			}
			if got := l.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestLinkedList_Erase tests the Erase method of the LinkedList
func TestLinkedList_Erase(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}

	t.Run("Test empty list", func(t *testing.T) {
		// Test empty list
		l.Erase(0)
	})

	t.Run("Test outrange index", func(t *testing.T) {
		l = GenerateRandomList(numberOfTests)
		// Test outrange index
		l.Erase(numberOfTests + 1)
	})

	t.Run("Test last index", func(t *testing.T) {
		l := LinkedList{}
		for i := 0; i < numberOfTests; i++ {
			l.PushBack(i)
		}
		// Test negative index
		l.Erase(-1)
		if l.Back() != numberOfTests-2 {
			t.Errorf("Erase() = %v, want %v", l.Back(), numberOfTests-2)
		}
		if l.Size() != numberOfTests-1 {
			t.Errorf("Size() = %v, want %v", l.Size(), numberOfTests-1)
		}
	})

	t.Run("Test correct index", func(t *testing.T) {
		l := LinkedList{}
		for i := 0; i < numberOfTests; i++ {
			l.PushBack(i)
		}

		for i := 0; i < numberOfTests; i++ {
			// get a random index to erase from the list between 0 and the size of the list
			index := rand.Intn(l.Size())
			l.Erase(index)
			if l.Size() != numberOfTests-i-1 {
				t.Errorf("Erase() = %v, want %v, index %v", l.Size(), numberOfTests-i-1, index)
				break
			}
		}
	})

	t.Run("Test negative index -9999", func(t *testing.T) {
		l := LinkedList{}
		l.Erase(-9999)
	})

	t.Run("Test negative index", func(t *testing.T) {
		l := LinkedList{}
		for i := 0; i < numberOfTests; i++ {
			l.PushBack(i)
		}
		for i := (numberOfTests - 1) * -1; i < 0; i++ {
			l.Erase(i * -1)
			if l.Size() != i*-1 {
				t.Errorf("Erase() = %v, want %v", l.Size(), i*-1)
			}
		}
	})
}

func TestLinkedList_Front(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	// Test empty list
	if l.Front() != 0 {
		t.Errorf("Front() = %v, want %v", l.Front(), 0)
	}

	for i := 0; i < numberOfTests; i++ {
		l.PushFront(i)
		if l.Front() != i {
			t.Errorf("Front() = %v, want %v", l.Front(), i)
		}
	}
}

func TestLinkedList_Insert(t *testing.T) {

	t.Run("Test empty list", func(t *testing.T) {
		l := LinkedList{}
		l.Insert(0, 1)
		if l.Size() != 1 {
			t.Errorf("Size() = %v, want %v", l.Size(), 1)
		}
	})
	t.Run("Test negative index", func(t *testing.T) {
		l := LinkedList{}
		l.Insert(-1, 1)
		if l.Size() != 1 {
			t.Errorf("Size() = %v, want %v", l.Size(), 0)
		}
	})

	t.Run("Test index 0", func(t *testing.T) {
		l := LinkedList{}
		l.Insert(0, 1)
		if l.Size() != 1 {
			t.Errorf("Size() = %v, want %v", l.Size(), 1)
		}
	})
	t.Run("Test index 1", func(t *testing.T) {
		l := NewLinkedList()
		l.PushBack(1)
		l.Insert(1, 1)
		if l.Size() != 2 {
			t.Errorf("Size() = %v, want %v", l.Size(), 2)
		}
	})

	t.Run("Test index 2", func(t *testing.T) {
		l := LinkedList{}
		l.PushBack(1)
		l.PushBack(2)
		l.Insert(1, 3)
		if l.Size() != 3 {
			t.Errorf("Size() = %v, want %v", l.Size(), 3)
		}

		if l.head.next.value != 3 {
			t.Errorf("Insert() = %v, want %v", l.head.next.value, 3)
		}

		if l.Back() != 2 {
			t.Errorf("Back() = %v, want %v", l.Back(), 2)
		}

		if l.Front() != 1 {
			t.Errorf("Front() = %v, want %v", l.Front(), 1)
		}
	})

}

// TestLinkedList_PopBack tests the PopBack method of the LinkedList
func TestLinkedList_PopBack(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	for i := 0; i < numberOfTests; i++ {
		l.PushBack(i)
	}

	for i := numberOfTests - 1; i >= 0; i-- {
		if l.Back() != i {
			t.Errorf("PopBack() = %v, want %v", l.Back(), i)
		}
		l.PopBack()
	}

	if !l.Empty() {
		t.Errorf("Empty() = %v, want %v", l.Empty(), true)
	}
}

// TestLinkedList_PopFront tests the PopFront method of the LinkedList
func TestLinkedList_PopFront(t *testing.T) {
	numberOfTests := 100
	t.Run("Test empty list", func(t *testing.T) {
		l := LinkedList{}
		l.PopFront()
		if !l.Empty() {
			t.Errorf("Empty() = %v, want %v", l.Empty(), true)
		}
	})
	t.Run("Test list with 1 element", func(t *testing.T) {
		l := LinkedList{}
		l.PushFront(1)
		l.PopFront()
		if !l.Empty() {
			t.Errorf("Empty() = %v, want %v", l.Empty(), true)
		}
	})
	t.Run("Test list with n elements", func(t *testing.T) {
		l := LinkedList{}
		for i := 0; i < numberOfTests; i++ {
			l.PushBack(i)
		}

		for i := 0; i < numberOfTests; i++ {
			if l.Front() != i {
				t.Errorf("PopFront() = %v, want %v", l.Front(), i)
			}
			l.PopFront()
		}

		if !l.Empty() {
			t.Errorf("Empty() = %v, want %v", l.Empty(), true)
		}
	})
}

// TestLinkedList_PushBack tests the PushBack method of the LinkedList
func TestLinkedList_PushBack(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	for i := 0; i < numberOfTests; i++ {
		l.PushBack(i)
		if l.Back() != i {
			t.Errorf("PushBack() = %v, want %v", l.Back(), i)
		}
	}

	if l.Size() != numberOfTests {
		t.Errorf("Size() = %v, want %v", l.Size(), numberOfTests)
	}
}

// TestLinkedList_PushFront tests the PushFront method of the LinkedList
func TestLinkedList_PushFront(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	for i := 0; i < numberOfTests; i++ {
		l.PushFront(i)
		if l.Front() != i {
			t.Errorf("PushFront() = %v, want %v", l.Front(), i)
		}
	}

	if l.Size() != numberOfTests {
		t.Errorf("Size() = %v, want %v", l.Size(), numberOfTests)
	}
}

// TestLinkedList_Size tests the size of the linked list
func TestLinkedList_Size(t *testing.T) {
	type fields struct {
		head *Node
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "test1",
			fields: fields{
				head: &Node{
					value: 1,
					next: &Node{
						value: 2,
						next: &Node{
							value: 3,
							next:  nil,
						},
					},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LinkedList{
				head: tt.fields.head,
			}
			if got := l.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestLinkedList_Size_Random tests the size of the linked list with random values
func TestLinkedList_Size_Random(t *testing.T) {
	numberOfNodes := []int{10, 100, 1000, 10000, 100000, 1000000}
	for _, n := range numberOfNodes {
		l := LinkedList{}
		for i := 0; i < n; i++ {
			l.PushFront(i)
		}
		if l.Size() != n {
			t.Errorf("Size() = %v, want %v", l.Size(), n)
		}
	}
}

func TestLinkedList_ValueAt(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	for i := 0; i < numberOfTests; i++ {
		l.PushBack(i)
	}

	for i := 0; i < numberOfTests; i++ {
		if l.ValueAt(i) != i {
			t.Errorf("ValueAt() = %v, want %v", l.ValueAt(i), i)
		}
	}
}

func TestLinkedList_ValueNFromEnd(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	for i := 0; i < numberOfTests; i++ {
		l.PushBack(i)
	}

	for i := 0; i < numberOfTests; i++ {
		if l.ValueNFromEnd(i) != numberOfTests-i-1 {
			t.Errorf("ValueNFromEnd() = %v, want %v", l.ValueNFromEnd(i), numberOfTests-i-1)
		}
	}
}

func TestNode_Value(t *testing.T) {
	numberOfTests := 100
	l := LinkedList{}
	for i := 0; i < numberOfTests; i++ {
		l.PushBack(i)
	}

	for i := 0; i < numberOfTests; i++ {
		if l.head.Value() != i {
			t.Errorf("ValueAt() = %v, want %v", l.ValueAt(i), i)
		}
		l.head = l.head.Next()
	}
}

func TestLinkedList_NewFromSlice(t *testing.T) {
	numberOfTests := 100
	s := make([]int, numberOfTests)
	for i := 0; i < numberOfTests; i++ {
		s[i] = i
	}

	l := NewFromSlice(s)

	for i := 0; i < numberOfTests; i++ {
		if l.ValueAt(i) != i {
			t.Errorf("ValueAt() = %v, want %v", l.ValueAt(i), i)
		}
	}
}
