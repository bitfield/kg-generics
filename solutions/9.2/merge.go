package merge

import "golang.org/x/exp/maps"

func Merge[K comparable, V any](ms ...map[K]V) map[K]V {
	result := map[K]V{}
	for _, m := range ms {
		maps.Copy(result, m)
	}
	return result
}
