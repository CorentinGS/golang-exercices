package queue

type Queue []interface{}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(values ...interface{}) {
	for _, v := range values {
		*q = append(*q, v)
	}
}

func (q *Queue) Dequeue() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	v := (*q)[0]
	(*q)[0] = nil // avoid memory leak

	*q = (*q)[1:]
	return v, true

}

func (q *Queue) Peek() (interface{}, bool) {
	if q.IsEmpty() {
		return nil, false
	}

	v := (*q)[0]
	return v, true
}

func (q *Queue) Size() int {
	return len(*q)
}

func (q *Queue) Clear() {
	*q = nil
}
