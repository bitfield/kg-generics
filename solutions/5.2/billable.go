package billable

type Channel[T any] struct {
	ch              chan T
	sends, receives int
}

func NewChannel[T any](length int) Channel[T] {
	return Channel[T]{
		ch: make(chan T, length),
	}
}

func (c *Channel[T]) Send(v T) {
	c.ch <- v
	c.sends++
}

func (c *Channel[T]) Receive() T {
	v := <-c.ch
	c.receives++
	return v
}

func (c Channel[T]) Sends() int {
	return c.sends
}

func (c Channel[T]) Receives() int {
	return c.receives
}
