package billable_test

import (
	"billable"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestSendReceive(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[int](1)
	want := 99
	c.Send(want)
	got := c.Receive()
	if !cmp.Equal(want, got) {
		t.Fatal(cmp.Diff(want, got))
	}
}

func TestSends(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[float64](3)
	c.Send(1.0)
	c.Send(2.0)
	c.Send(3.0)
	want := 3
	got := c.Sends()
	if want != got {
		t.Fatalf("want %d sends, got %d", want, got)
	}
}

func TestReceives(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[struct{}](1)
	c.Send(struct{}{})
	_ = c.Receive()
	c.Send(struct{}{})
	_ = c.Receive()
	want := 2
	got := c.Receives()
	if want != got {
		t.Fatalf("want 2 receives, got %d", got)
	}
}
