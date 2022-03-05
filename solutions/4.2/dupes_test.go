package dupes_test

import (
	"dupes"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDupesIsTrueWhenInputContainsNonConsecutiveDuplicates(t *testing.T) {
	t.Parallel()
	s := []int{1, 2, 3, 1, 5}
	want := true
	got := dupes.Dupes(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestDupesIsFalseWhenInputContainsNoDuplicates(t *testing.T) {
	t.Parallel()
	s := []string{"a"}
	want := false
	got := dupes.Dupes(s)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
