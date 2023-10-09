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
