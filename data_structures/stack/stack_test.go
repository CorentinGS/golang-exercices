package stack

import "testing"

func TestStack_IsEmpty(t *testing.T) {
	t.Run("empty stack", func(t *testing.T) {
		s := Stack{}
		if !s.IsEmpty() {
			t.Errorf("expected stack to be empty")
		}
	})

	t.Run("non-empty stack", func(t *testing.T) {
		s := Stack{}
		s.Push(1)
		if s.IsEmpty() {
			t.Errorf("expected stack to be non-empty")
		}
	})
}

func TestStack_Push(t *testing.T) {
	t.Run("push to empty stack", func(t *testing.T) {
		s := Stack{}
		s.Push(1)
		s.Push(2)
		s.Push(3)

		if s.Size() != 3 {
			t.Errorf("expected stack size to be 3, got %d", s.Size())
		}
	})

	t.Run("push to non-empty stack", func(t *testing.T) {
		s := Stack{4, 5, 6}
		s.Push(1)
		s.Push(2)
		s.Push(3)
		if s.Size() != 6 {
			t.Errorf("expected stack size to be 6, got %d", s.Size())
		}
	})

	t.Run("push multiple values", func(t *testing.T) {
		s := Stack{}
		s.Push(1, 2, 3)
		if s.Size() != 3 {
			t.Errorf("expected stack size to be 3, got %d", s.Size())
		}
	})

	t.Run("push multiple values to non-empty stack", func(t *testing.T) {
		s := Stack{4, 5, 6}
		s.Push(1, 2, 3)
		if s.Size() != 6 {
			t.Errorf("expected stack size to be 6, got %d", s.Size())
		}
	})

	t.Run("push multiple values with different types", func(t *testing.T) {
		s := Stack{4, 5, 6}
		s.Push(1, "2", 4.5435)
		if s.Size() != 6 {
			t.Errorf("expected stack size to be 6, got %d", s.Size())
		}
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("pop from empty stack", func(t *testing.T) {
		s := Stack{}
		if _, ok := s.Pop(); ok {
			t.Errorf("expected pop to fail")
		}
	})

	t.Run("pop from non-empty stack", func(t *testing.T) {
		s := Stack{1, 2, 3}
		if v, ok := s.Pop(); !ok || v != 3 {
			t.Errorf("expected pop to return 3, got %v", v)
		}
	})
}

func TestStack_Peek(t *testing.T) {
	t.Run("peek from empty stack", func(t *testing.T) {
		s := Stack{}
		if _, ok := s.Peek(); ok {
			t.Errorf("expected peek to fail")
		}
	})

	t.Run("peek from non-empty stack", func(t *testing.T) {
		s := Stack{1, 2, 3}
		if v, ok := s.Peek(); !ok || v != 3 {
			t.Errorf("expected peek to return 3, got %v", v)
		}
	})
}

func TestStack_Size(t *testing.T) {
	t.Run("size of empty stack", func(t *testing.T) {
		s := Stack{}
		if s.Size() != 0 {
			t.Errorf("expected stack size to be 0, got %d", s.Size())
		}
	})

	t.Run("size of non-empty stack", func(t *testing.T) {
		s := Stack{1, 2, 3}
		if s.Size() != 3 {
			t.Errorf("expected stack size to be 3, got %d", s.Size())
		}
	})
}

func TestStack_Clear(t *testing.T) {
	t.Run("clear empty stack", func(t *testing.T) {
		s := Stack{}
		s.Clear()
		if s.Size() != 0 {
			t.Errorf("expected stack size to be 0, got %d", s.Size())
		}
	})

	t.Run("clear non-empty stack", func(t *testing.T) {
		s := Stack{1, 2, 3}
		s.Clear()
		if s.Size() != 0 {
			t.Errorf("expected stack size to be 0, got %d", s.Size())
		}
	})

	t.Run("clear stack with multiple values", func(t *testing.T) {
		s := Stack{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		s.Clear()
		if s.Size() != 0 {
			t.Errorf("expected stack size to be 0, got %d", s.Size())
		}

		s.Push(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
		if s.Size() != 10 {
			t.Errorf("expected stack size to be 10, got %d", s.Size())
		}
	})
}

func Fuzz_Stack(f *testing.F) {
	f.Add(5, -52.1454, "dsfxd")
	f.Add(-4, -23.34, "#{`~#[{|")
	f.Add(5643, 423.423, "test")
	f.Fuzz(func(t *testing.T, i int, float float64, str string) {
		s := Stack{}
		s.Push(i, float, str)
		s.Clear()
	})
}
