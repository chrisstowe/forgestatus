package common

import (
	"math/rand"
	"strconv"
)

const (
	mockTime = "2019-02-11T18:27:02.566923-08:00"
	mockID   = "9m4e2mr0ui3e8a215n4g"
)

// MockSystemMetric returns a simulated system metric.
// The upper bounds (non-inclusive) is specified by n.
func MockSystemMetric(n int) string {
	return strconv.Itoa(rand.Intn(n))
}
