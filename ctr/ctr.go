package ctr

import (
	"sync/atomic"
)

// Counter type is alias for int64 with attached methods
type Counter int64

// Increment adds 1 to couter using atomic
func (c *Counter) Increment() {
	atomic.AddInt64((*int64)(c), 1)
}

// Decrement adds -1 to couter using atomic
func (c *Counter) Decrement() {
	atomic.AddInt64((*int64)(c), -1)
}

// Reset sets counter to 0 using atomic
func (c *Counter) Reset() {
	atomic.StoreInt64((*int64)(c), 0)
}
