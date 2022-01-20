package product

import "constraints"

type Number interface {
	constraints.Integer | constraints.Float | constraints.Complex
}

func Product[T Number](x, y T) T {
	return x * y
}
