package group_test

import (
	"group"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLenOfGroupIs2WhenItContains2Elements(t *testing.T) {
	t.Parallel()
	g := group.Group[int]{1, 2}
	want := 2
	got := group.Len(g)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
