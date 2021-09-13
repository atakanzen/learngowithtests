package synk

import "sync"

type Counter struct {
	mu  sync.Mutex
	val int
}

// Returns a new pointer to the counter
func NewCounter() *Counter {
	return &Counter{}
}

// Increments the counter's value
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

// Returns the counter's value
func (c *Counter) Value() int {
	return c.val
}
