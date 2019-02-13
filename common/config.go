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
	MaxTaskQueueSize   int64
	MaxResultQueueSize int64
}

// EnvConfig holds required and optional environment parameters.
var EnvConfig = newConfig()

func newConfig() Config {
	port := os.Getenv("PORT")
	redisURL := os.Getenv("REDIS_URL")
	workerID := os.Getenv("WORKER_ID")

	maxTaskQueue := os.Getenv("MAX_TASK_QUEUE_SIZE")
	maxTaskQueueSize, err := strconv.ParseInt(maxTaskQueue, 10, 64)
	if err != nil {
		maxTaskQueueSize = 100
	}

	maxResultQueue := os.Getenv("MAX_RESULT_QUEUE_SIZE")
	maxResultQueueSize, err := strconv.ParseInt(maxResultQueue, 10, 64)
	if err != nil {
		maxResultQueueSize = 100
	}

	return Config{
		Port:               port,
		RedisURL:           redisURL,
		WorkerID:           workerID,
		MaxTaskQueueSize:   maxTaskQueueSize,
		MaxResultQueueSize: maxResultQueueSize,
	}
}
