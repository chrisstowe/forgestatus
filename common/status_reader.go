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

func getResults(sr *statusReader, t TaskType) ([]Result, error) {
	resultQueue := ResultQueuePrefix + string(t)

	// Try and grab the maximum number of results.
	serializedResults, err := sr.client.LRange(resultQueue,
		0, EnvConfig.MaxResultQueueSize-1).Result()
	if err != nil {
		return nil, err
	}

	results := make([]Result, 0, len(serializedResults))
	for _, s := range serializedResults {
		result, err := DeserializeResult(s)
		if err != nil {
			return nil, err
		}

		results = append(results, *result)
	}

	return results, nil
}

func (sr *statusReader) GetStatus() (*Status, error) {
	tasksScheduled, err := sr.client.Get(TasksScheduledCounter).Result()
	if err != nil {
		return nil, err
	}

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

	healthy, err := getResults(sr, GetHealthy)
	if err != nil {
		return nil, err
	}

	ready, err := getResults(sr, GetReady)
	if err != nil {
		return nil, err
	}

	memoryUsed, err := getResults(sr, GetMemoryUsed)
	if err != nil {
		return nil, err
	}

	cpuUsed, err := getResults(sr, GetCPUUsed)
	if err != nil {
		return nil, err
	}

	diskUsed, err := getResults(sr, GetDiskUsed)
	if err != nil {
		return nil, err
	}

	procsRunning, err := getResults(sr, GetProcsRunning)
	if err != nil {
		return nil, err
	}

	diskIO, err := getResults(sr, GetDiskIO)
	if err != nil {
		return nil, err
	}

	networkTraffic, err := getResults(sr, GetNetworkTraffic)
	if err != nil {
		return nil, err
	}

	status := &Status{
		TasksScheduled: tasksScheduled,
		TasksProcessed: tasksProcessed,
		Healthy:        healthy,
		Ready:          ready,
		MemoryUsed:     memoryUsed,
		CPUUsed:        cpuUsed,
		DiskUsed:       diskUsed,
		ProcsRunning:   procsRunning,
		DiskIO:         diskIO,
		NetworkTraffic: networkTraffic,
	}

	return status, nil
}
