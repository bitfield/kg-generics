package stringy_test

import (
	"bytes"
	"stringy"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type greeting struct{}

func (greeting) String() string {
	return "Howdy!"
}

func TestStringifyTo(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	stringy.StringifyTo[greeting](buf, greeting{})
	want := "Howdy!\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
