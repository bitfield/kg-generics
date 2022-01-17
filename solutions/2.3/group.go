package group

type Group[E any] []E

func Len[E any](s []E) int {
	return len(s)
}
