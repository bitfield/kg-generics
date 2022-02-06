package product

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Product[T Number](x, y T) T {
	return x * y
}
