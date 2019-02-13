package common

import (
	"strconv"

	"github.com/go-redis/redis"
)

// StatusReader reads the current status from the database.
type StatusReader interface {
	GetStatus() (*Status, error)
}

type statusReader struct {
	client *redis.Client
}

// NewStatusReader creates a new StatusReader.
func NewStatusReader(redisURL string) StatusReader {
	c := redis.NewClient(&redis.Options{Addr: redisURL})
	return &statusReader{client: c}
}

func (sr *statusReader) GetStatus() (*Status, error) {
	// memQueue := ResultQueuePrefix + string(GetMemoryUsed)
	// memoryUsed, err := sr.client.LRange(memQueue, 0, EnvConfig.MaxResultQueueSize-1).Result()
	// if err != nil {
	// 	return nil, err
	// }

	// The reader does not know how many workers there are.
	// This is intentional (since workers could scale up or down).
	// The strategy here is to read from worker counts until failure.
	// This decoupling could likely happen through persistence in the future.
	tasksProcessed := make([]string, 0, 10)
	workerID := 1
	for {
		processedCounter := TasksProcessedCounterPrefix + strconv.Itoa(workerID)
		pc, err := sr.client.Get(processedCounter).Result()
		if err != nil {
			break
		}

		tasksProcessed = append(tasksProcessed, pc)
		workerID++
	}

	tasksScheduled, err := sr.client.Get(TasksScheduledCounter).Result()
	if err != nil {
		return nil, err
	}

	status := &Status{
		MemoryUsed:     []string{"0"},
		CPUUsed:        []string{"0"},
		DiskUsed:       []string{"0"},
		ProcsRunning:   []string{"0"},
		TasksProcessed: tasksProcessed,
		TasksScheduled: tasksScheduled,
	}

	return status, nil
}
