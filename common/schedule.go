package common

import "time"

// Schedule calls a function on a periodic interval.
func Schedule(f func(), interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			f()
		}
	}()
}
