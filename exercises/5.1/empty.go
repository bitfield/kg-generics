package empty

type Sequence[E any] []E

func (s Sequence[E]) Empty() bool {
	return len(s) == 0
}
