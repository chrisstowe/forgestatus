package common

import (
	"os"
	"strconv"
)

// Config holds required and optional system parameters.
type Config struct {
	Port               string
	RedisURL           string
	WorkerID           string
	WorkerCount        int
	MaxTaskQueueSize   int64
	MaxResultQueueSize int64
}

// EnvConfig holds required and optional environment parameters.
var EnvConfig = newConfig()

func newConfig() Config {
	port := os.Getenv("PORT")
	redisURL := os.Getenv("REDIS_URL")
	workerID := os.Getenv("WORKER_ID")

	wc := os.Getenv("WORKER_COUNT")
	workerCount, err := strconv.Atoi(wc)
	if err != nil || workerCount < 1 {
		workerCount = 1
	}

	mt := os.Getenv("MAX_TASK_QUEUE_SIZE")
	maxTaskQueueSize, err := strconv.ParseInt(mt, 10, 64)
	if err != nil || maxTaskQueueSize < 1 {
		maxTaskQueueSize = 100
	}

	mr := os.Getenv("MAX_RESULT_QUEUE_SIZE")
	maxResultQueueSize, err := strconv.ParseInt(mr, 10, 64)
	if err != nil || maxResultQueueSize < 1 {
		maxResultQueueSize = 100
	}

	return Config{
		Port:               port,
		RedisURL:           redisURL,
		WorkerID:           workerID,
		WorkerCount:        workerCount,
		MaxTaskQueueSize:   maxTaskQueueSize,
		MaxResultQueueSize: maxResultQueueSize,
	}
}
