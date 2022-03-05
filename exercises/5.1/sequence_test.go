package sequence_test

import (
	"sequence"
	"testing"
)

func TestEmptyIsTrueForEmptySequence(t *testing.T) {
	t.Parallel()
	s := sequence.Sequence[int]{}
	got := s.Empty()
	if !got {
		t.Fatal("false for empty sequence")
	}
}

func TestEmptyIsFalseForNonEmptySequence(t *testing.T) {
	t.Parallel()
	s := sequence.Sequence[string]{"a", "b", "c"}
	got := s.Empty()
	if got {
		t.Fatal("true for non-empty sequence")
	}
}
