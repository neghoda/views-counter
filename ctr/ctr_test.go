package ctr

import (
	"runtime"
	"sync"
	"testing"
)

// Not necessary as Increment and Decrement only do a call to atomic
// Unsync counter (almost) always fails in this tests with 500 routines
const routines int = 500

func TestIncrement(t *testing.T) {
	var c Counter = 0
	var wg sync.WaitGroup
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			runtime.Gosched()
			c.Increment()
			wg.Done()
		}()
	}
	wg.Wait()

	t.Run("Increment operation should be atomic", func(t *testing.T) {
		if int(c) != routines {
			t.Errorf("Expected counter value - \"%v\" to be \"%v\"", c, routines)
		}
	})
}

func TestDecrement(t *testing.T) {
	var c Counter = 0
	var wg sync.WaitGroup
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			runtime.Gosched()
			c.Decrement()
			wg.Done()
		}()
	}
	wg.Wait()

	t.Run("Decrement operation should be atomic", func(t *testing.T) {
		if int(c) != -routines {
			t.Errorf("Expected counter value - \"%v\" to be \"%v\"", c, -routines)
		}
	})
}
