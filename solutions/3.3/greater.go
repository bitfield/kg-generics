package greater

func IsGreater[T interface{ Greater(T) bool }](x, y T) bool {
	return x.Greater(y)
}
