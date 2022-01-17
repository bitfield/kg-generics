package group_test

import (
	"group"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGroup(t *testing.T) {
	t.Parallel()
	got := group.Group[string]{}
	got = append(got, "hello")
	got = append(got, "world")
	want := group.Group[string]{"hello", "world"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
