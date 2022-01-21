package stack

type Stack[E any] struct {
	data []E
}

func (s Stack[E]) Len() int {
	return len(s.data)
}

func (s *Stack[E]) Push(vals ...E) {
	s.data = append(s.data, vals...)
}

func (s *Stack[E]) Pop() (v E, ok bool) {
	if len(s.data) == 0 {
		return v, false
	}
	v = s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v, true
}
