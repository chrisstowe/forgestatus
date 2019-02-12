package common

import (
	"os"
	"strconv"
)

// Config holds required and optional system parameters.
type Config struct {
	Port             string
	RedisURL         string
	WorkerID         string
	MaxTaskQueueSize int64
}

// EnvConfig holds required and optional environment parameters.
var EnvConfig = newConfig()

func newConfig() Config {
	port := os.Getenv("PORT")
	redisURL := os.Getenv("REDIS_URL")
	workerID := os.Getenv("WORKER_ID")

	s := os.Getenv("MAX_TASK_QUEUE_SIZE")
	maxTaskQueueSize, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		maxTaskQueueSize = 100
	}

	return Config{
		Port:             port,
		RedisURL:         redisURL,
		WorkerID:         workerID,
		MaxTaskQueueSize: maxTaskQueueSize,
	}
}
