package stack

type Stack []interface{}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(values ...interface{}) {
	for _, v := range values {
		*s = append(*s, v)
	}
}

func (s *Stack) Pop() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	index := len(*s) - 1
	v := (*s)[index]
	(*s)[index] = nil // avoid memory leak
	*s = (*s)[:index]
	return v, true
}

func (s *Stack) Peek() (interface{}, bool) {
	if s.IsEmpty() {
		return nil, false
	}

	index := len(*s) - 1
	v := (*s)[index]
	return v, true
}

func (s *Stack) Size() int {
	return len(*s)
}

func (s *Stack) Clear() {
	*s = nil
}
