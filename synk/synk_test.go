package synk_test

import (
	"learngowithtests/synk"
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing counter 3 times leaves the value at 3", func(t *testing.T) {
		counter := synk.NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := synk.NewCounter()

		var wg sync.WaitGroup
		// Set the number of goroutines to wait for
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				// Mark the current goroutine as Done
				wg.Done()
			}()
		}

		// Wait for all goroutines to finish
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *synk.Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
