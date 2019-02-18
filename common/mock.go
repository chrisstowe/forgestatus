package common

import (
	"math/rand"
	"strconv"
)

const (
	mockTime = "2019-02-11T18:27:02.566923-08:00"
	mockID   = "9m4e2mr0ui3e8a215n4g"
)

var taskTypes = []TaskType{
	GetHealthy,
	GetReady,
	GetMemoryUsed,
	GetCPUUsed,
	GetDiskUsed,
	GetProcsRunning,
	GetDiskIO,
	GetNetworkTraffic,
}

// MockHealthyReady returns a simulated healthy/ready status.
func MockHealthyReady() string {
	return strconv.FormatBool(rand.Intn(15) > 0)
}

// MockSystemValue returns a simulated system value.
// The upper bounds (non-inclusive) is specified by n.
func MockSystemValue(n int) string {
	return strconv.Itoa(rand.Intn(n))
}
