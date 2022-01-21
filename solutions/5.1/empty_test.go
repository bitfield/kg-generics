package empty_test

import (
	"empty"
	"testing"
)

func TestEmptyTrue(t *testing.T) {
	t.Parallel()
	s := empty.Sequence[int]{}
	got := s.Empty()
	if !got {
		t.Fatal("false for empty sequence")
	}
}

func TestEmptyFalse(t *testing.T) {
	t.Parallel()
	s := empty.Sequence[string]{"a", "b", "c"}
	got := s.Empty()
	if got {
		t.Fatal("true for non-empty sequence")
	}
}
