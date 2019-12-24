package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter 3 times laves it at 3", func(t *testing.T) {
		c := NewCounter()
		c.Inc()
		c.Inc()
		c.Inc()

		assertCounter(t, c, 3)
	})

	t.Run("It runs safely concurrently", func(t *testing.T) {
		want := 1000
		c := NewCounter()

		var wg sync.WaitGroup
		wg.Add(want)

		for i := 0; i < want; i++ {
			go func(w *sync.WaitGroup) {
				c.Inc()
				w.Done()
			}(&wg)
		}

		wg.Wait()
		assertCounter(t, c, want)
	})
}

func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d want %d", got.Value(), want)
	}
}
