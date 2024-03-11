package counter

import (
	"context"
	"sync"
)

// Counter is structure with 3 fields.
type Counter struct {
	value     int
	limit     int
	increment chan int
}

// NewCounter is the constructor.
func NewCounter(limit int) *Counter {
	c := Counter{
		limit:     limit,
		value:     0,
		increment: make(chan int, limit),
	}
	return &c
}

// Add is the method for sending message to channel. If the counter is greater than or equal max value, sending
// to channel stops.
func (c *Counter) Add(amount int, ctx context.Context, cancel context.CancelFunc) {
	if c.value >= c.limit {
		ctx.Done()
		cancel()
	}
	c.increment <- amount
}

// Increment is the method that increment counter while it is less than max value.
func (c *Counter) Increment(wg *sync.WaitGroup, cancel context.CancelFunc) {
	defer wg.Done()
	for step := range c.increment {
		if c.value < c.limit {
			c.value += step
		} else {
			cancel()
		}
	}
}

// Value is getter for field value in Counter struct.
func (c *Counter) Value() int {
	return c.value
}

// CloseChanel is method for closing the channel.
func (c *Counter) CloseChanel() {
	close(c.increment)
}
