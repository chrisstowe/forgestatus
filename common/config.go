package common

import (
	"os"
	"strconv"
)

// Config holds required and optional system parameters.
type Config struct {
	Env                string
	Port               string
	RedisURL           string
	WorkerID           int
	WorkerCount        int
	MaxTaskQueueSize   int64
	MaxResultQueueSize int64
}

// EnvConfig holds required and optional environment parameters.
var EnvConfig = newConfig()

func newConfig() Config {
	env := os.Getenv("ENV")
	port := os.Getenv("PORT")
	redisURL := os.Getenv("REDIS_URL")

	wi := os.Getenv("WORKER_ID")
	workerID, err := strconv.Atoi(wi)
	if err != nil || workerID < 1 {
		workerID = 1
	}

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
		Env:                env,
		Port:               port,
		RedisURL:           redisURL,
		WorkerID:           workerID,
		WorkerCount:        workerCount,
		MaxTaskQueueSize:   maxTaskQueueSize,
		MaxResultQueueSize: maxResultQueueSize,
	}
}
