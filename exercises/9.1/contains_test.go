package contains_test

import (
	"contains"
	"testing"
	"unicode"
)

func positive(p int) bool {
	return p > 0
}

func TestContainsFunc_IsTrueForPositiveOnInputContainingPositiveInts(t *testing.T) {
	t.Parallel()
	input := []int{-2, 0, 1, -1, 5}
	got := contains.ContainsFunc(input, positive)
	if !got {
		t.Fatalf("%v: want true for 'contains positive', got false", input)
	}
}

func TestContainsFunc_IsFalseForIsUpperOnInputContainingNoUppercaseRunes(t *testing.T) {
	t.Parallel()
	input := []rune("hello")
	got := contains.ContainsFunc(input, unicode.IsUpper)
	if got {
		t.Fatalf("%q: want false for 'contains uppercase', got true", input)
	}
}
