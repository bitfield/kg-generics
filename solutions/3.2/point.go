package point

type Point struct {
	X, Y int
}

type OnlyPoint interface {
	Point
}

func GetX[T OnlyPoint](p T) int {
	return p.X
}

func GetY[T OnlyPoint](p T) int {
	return p.Y
}
