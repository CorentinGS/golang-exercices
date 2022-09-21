package queue

import "testing"

func TestQueue_IsEmpty(t *testing.T) {
	t.Run("empty queue", func(t *testing.T) {
		q := Queue{}
		if !q.IsEmpty() {
			t.Errorf("expected queue to be empty")
		}
	})

	t.Run("non-empty queue", func(t *testing.T) {
		q := Queue{}
		q.Enqueue(1)
		if q.IsEmpty() {
			t.Errorf("expected queue to be non-empty")
		}
	})
}

func TestQueue_Enqueue(t *testing.T) {
	t.Run("enqueue to empty queue", func(t *testing.T) {
		q := Queue{}
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		if q.Size() != 3 {
			t.Errorf("expected queue size to be 3, got %d", q.Size())
		}
	})

	t.Run("enqueue to non-empty queue", func(t *testing.T) {
		q := Queue{4, 5, 6}
		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)
		if q.Size() != 6 {
			t.Errorf("expected queue size to be 6, got %d", q.Size())
		}
	})

	t.Run("enqueue multiple values", func(t *testing.T) {
		q := Queue{}
		q.Enqueue(1, 2, 3)
		if q.Size() != 3 {
			t.Errorf("expected queue size to be 3, got %d", q.Size())
		}
	})

	t.Run("enqueue multiple values to non-empty queue", func(t *testing.T) {
		q := Queue{4, 5, 6}
		q.Enqueue(1.342, -2, "3")
		if q.Size() != 6 {
			t.Errorf("expected queue size to be 6, got %d", q.Size())
		}
	})
}

func TestQueue_Dequeue(t *testing.T) {
	t.Run("dequeue from empty queue", func(t *testing.T) {
		q := Queue{}
		if _, ok := q.Dequeue(); ok {
			t.Errorf("expected error when dequeueing from empty queue")
		}
	})

	t.Run("dequeue from non-empty queue", func(t *testing.T) {
		q := Queue{1, 2, 3}
		if v, ok := q.Dequeue(); !ok {
			t.Errorf("expected no error when dequeueing from non-empty queue")
		} else if v != 1 {
			t.Errorf("expected value to be 1, got %v", v)
		}
	})
}

func TestQueue_Size(t *testing.T) {
	t.Run("size of empty queue", func(t *testing.T) {
		q := Queue{}
		if q.Size() != 0 {
			t.Errorf("expected size to be 0, got %d", q.Size())
		}
	})

	t.Run("size of non-empty queue", func(t *testing.T) {
		q := Queue{1, 2, 3}
		if q.Size() != 3 {
			t.Errorf("expected size to be 3, got %d", q.Size())
		}
	})
}

func TestQueue_Clear(t *testing.T) {
	t.Run("clear empty queue", func(t *testing.T) {
		q := Queue{}
		q.Clear()
		if q.Size() != 0 {
			t.Errorf("expected size to be 0, got %d", q.Size())
		}
	})

	t.Run("clear non-empty queue", func(t *testing.T) {
		q := Queue{1, 2, 3}
		q.Clear()
		if q.Size() != 0 {
			t.Errorf("expected size to be 0, got %d", q.Size())
		}
	})
}

func TestQueue_Peek(t *testing.T) {
	t.Run("peek empty queue", func(t *testing.T) {
		q := Queue{}
		if _, ok := q.Peek(); ok {
			t.Errorf("expected error when peeking empty queue")
		}
	})

	t.Run("peek non-empty queue", func(t *testing.T) {
		q := Queue{1, 2, 3}
		if v, ok := q.Peek(); !ok {
			t.Errorf("expected no error when peeking non-empty queue")
		} else if v != 1 {
			t.Errorf("expected value to be 1, got %v", v)
		}
	})
}
