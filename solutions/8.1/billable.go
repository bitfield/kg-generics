package billable

import "sync"

type Channel[T any] struct {
	lock            *sync.RWMutex
	ch              chan T
	sends, receives int
}

func NewChannel[T any](length int) *Channel[T] {
	return &Channel[T]{
		lock: &sync.RWMutex{},
		ch:   make(chan T, length),
	}
}

func (c *Channel[T]) Send(v T) {
	c.ch <- v
	c.lock.Lock()
	defer c.lock.Unlock()
	c.sends++
}

func (c *Channel[T]) Receive() T {
	v := <-c.ch
	c.lock.Lock()
	defer c.lock.Unlock()
	c.receives++
	return v
}

func (c *Channel[T]) Sends() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.sends
}

func (c *Channel[T]) Receives() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.receives
}
