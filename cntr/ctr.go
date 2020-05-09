package ctr

import (
	"sync/atomic"
)

type Counter int64

func (c *Counter) Increment() {
	atomic.AddInt64((*int64)(c), 1)
}

func (c *Counter) Decrement() {
	atomic.AddInt64((*int64)(c), -1)
}

func (c *Counter) Reset() {
	atomic.StoreInt64((*int64)(c), 0)
}
