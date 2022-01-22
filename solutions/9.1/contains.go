package contains

import "golang.org/x/exp/slices"

func ContainsFunc[E any](s []E, f func(E) bool) bool {
	return slices.IndexFunc(s, f) >= 0
}
