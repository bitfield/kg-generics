package print_test

import (
	"bytes"
	"print"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestPrintAnythingTo_PrintsInputToSuppliedWriter(t *testing.T) {
	t.Parallel()
	buf := &bytes.Buffer{}
	print.PrintAnythingTo[string](buf, "Hello, world")
	want := "Hello, world\n"
	got := buf.String()
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
