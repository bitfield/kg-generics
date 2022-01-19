package intish

type Intish interface {
	~int
}

func IsPositive[T Intish](v T) bool {
	return v > 0
}
