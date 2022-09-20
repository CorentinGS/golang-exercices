package linkedList

import (
	"math/rand"
	"testing"
)

func GenerateRandomDoubleList(n int) DoubleLinkedList {
	l := DoubleLinkedList{}
	for i := 0; i < n; i++ {
		i := rand.Intn(n)
		l.PushBack(i)
	}
	return l
}

func TestDoubleLinkedList_Back(t *testing.T) {
	numberOfTests := 100
	l := DoubleLinkedList{}
	// Test empty list
	if l.Back() != 0 {
		t.Errorf("Back() = %v, want %v", l.Back(), 0)
	}
	l = GenerateRandomDoubleList(numberOfTests)

	for i := 0; i < numberOfTests; i++ {
		l.PushBack(i)
		if l.Back() != i {
			t.Errorf("Back() = %v, want %v", l.Back(), i)
		}
	}
}

func TestDoubleLinkedList_Empty(t *testing.T) {
	type fields struct {
		head *DoubleNode
		tail *DoubleNode
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
				head: &DoubleNode{
					value: 1,
					next:  nil,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &DoubleLinkedList{
				head: tt.fields.head,
				tail: tt.fields.tail,
			}
			if got := l.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDoubleLinkedList_Erase(t *testing.T) {
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
		l := DoubleLinkedList{}
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
		l := DoubleLinkedList{}
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
		l := DoubleLinkedList{}
		l.Erase(-9999)
	})

	t.Run("Test negative index", func(t *testing.T) {
		l := DoubleLinkedList{}
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

func TestDoubleLinkedList_Front(t *testing.T) {
	numberOfTests := 100
	l := DoubleLinkedList{}
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

func TestDoubleLinkedList_Insert(t *testing.T) {

	t.Run("Test empty list", func(t *testing.T) {
		l := DoubleLinkedList{}
		l.Insert(0, 1)
		if l.Size() != 1 {
			t.Errorf("Size() = %v, want %v", l.Size(), 1)
		}
	})
	t.Run("Test negative index", func(t *testing.T) {
		l := DoubleLinkedList{}
		l.Insert(-1, 1)
		if l.Size() != 1 {
			t.Errorf("Size() = %v, want %v", l.Size(), 0)
		}
	})

	t.Run("Test index 0", func(t *testing.T) {
		l := DoubleLinkedList{}
		l.Insert(0, 1)
		if l.Size() != 1 {
			t.Errorf("Size() = %v, want %v", l.Size(), 1)
		}
	})
	t.Run("Test index 1", func(t *testing.T) {
		l := DoubleLinkedList{}
		l.PushBack(1)
		l.Insert(1, 1)
		if l.Size() != 2 {
			t.Errorf("Size() = %v, want %v", l.Size(), 2)
		}
	})

	t.Run("Test index 2", func(t *testing.T) {
		l := DoubleLinkedList{}
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

func TestDoubleLinkedList_PopBack(t *testing.T) {
	numberOfTests := 100
	l := DoubleLinkedList{}
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

func TestDoubleLinkedList_PopFront(t *testing.T) {
	numberOfTests := 100
	t.Run("Test empty list", func(t *testing.T) {
		l := DoubleLinkedList{}
		l.PopFront()
		if !l.Empty() {
			t.Errorf("Empty() = %v, want %v", l.Empty(), true)
		}
	})
	t.Run("Test list with 1 element", func(t *testing.T) {
		l := DoubleLinkedList{}
		l.PushFront(1)
		l.PopFront()
		if !l.Empty() {
			t.Errorf("Empty() = %v, want %v", l.Empty(), true)
		}
	})
	t.Run("Test list with n elements", func(t *testing.T) {
		l := DoubleLinkedList{}
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

func TestDoubleLinkedList_PushBack(t *testing.T) {
	numberOfTests := 100
	l := DoubleLinkedList{}
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

func TestDoubleLinkedList_PushFront(t *testing.T) {
	numberOfTests := 100
	l := DoubleLinkedList{}
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
