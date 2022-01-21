package billable_test

import (
	"billable"
	"sync"
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

func TestConcurrencySafety(t *testing.T) {
	t.Parallel()
	c := billable.NewChannel[string](10)
	want := 100
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i < want; i++ {
			c.Send("hello")
			_ = c.Receives()
		}
		wg.Done()
	}()
	for i := 0; i < want; i++ {
		_ = c.Receive()
		_ = c.Sends()
	}
	wg.Wait()
	got := c.Sends()
	if want != got {
		t.Errorf("want %d sends, got %d", want, got)
	}
	got = c.Receives()
	if want != got {
		t.Errorf("want %d receives, got %d", want, got)
	}
}
