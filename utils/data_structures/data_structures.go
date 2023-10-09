package data_structures

type Stack struct {
	elements []interface{}
}

func (s *Stack) Push(element interface{}) {
	s.elements = append(s.elements, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.elements) == 0 {
		return nil
	}

	topIndex := len(s.elements) - 1
	top := s.elements[topIndex]
	s.elements = s.elements[:topIndex]
	return top
}

func (s *Stack) Peek() interface{} {
	if len(s.elements) == 0 {
		return nil
	}

	return s.elements[len(s.elements)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack) Size() int {
	return len(s.elements)
}

type Queue struct {
	elements []interface{}
}

func (q *Queue) Push(element interface{}) {
	q.elements = append(q.elements, element)
}

func (q *Queue) Pop() interface{} {
	if len(q.elements) == 0 {
		return nil
	}

	front := q.elements[0]
	q.elements = q.elements[1:]
	return front
}

func (q *Queue) Peek() interface{} {
	if len(q.elements) == 0 {
		return nil
	}

	return q.elements[0]
}

func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}

func (q *Queue) Size() int {
	return len(q.elements)
}
