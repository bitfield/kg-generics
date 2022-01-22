package compose

func Compose[T, U, V any](f func(U) T, g func(V) U, v V) T {
	return f(g(v))
}
