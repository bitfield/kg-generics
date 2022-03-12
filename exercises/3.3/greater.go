package greater

func IsGreater[T /* Your constraint goes here! */](x, y T) bool {
	return x.Greater(y)
}
