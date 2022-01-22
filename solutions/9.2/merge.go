package merge

import "golang.org/x/exp/maps"

func Merge[M ~map[K]V, K comparable, V any](ms ...M) M {
	result := M{}
	for _, m := range ms {
		maps.Copy[map[K]V](result, m)
	}
	return result
}
