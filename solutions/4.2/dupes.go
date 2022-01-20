package dupes

func Dupes[E comparable](s []E) bool {
	seen := map[E]bool{}
	for _, v := range s {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
