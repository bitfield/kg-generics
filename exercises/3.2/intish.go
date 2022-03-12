package intish

// Your Intish interface goes here!

func IsPositive[T Intish](v T) bool {
	return v > 0
}
