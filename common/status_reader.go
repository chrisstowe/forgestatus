package common

import "github.com/go-redis/redis"

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
	// s, err := SerializeTask(task)
	// if err != nil {
	// 	return nil, err
	// }

	// err = ts.client.LPush(PendingQueue, s).Err()
	// if err != nil {
	// 	return nil, err
	// }

	// // Prevent too many tasks from building up.
	// // In a real system, the number of workers should probably be scaled up.
	// // This is an O(1) operation (since the worst case is always removing 1).
	// err = ts.client.LTrim(PendingQueue, 0, EnvConfig.MaxTaskQueueSize-1).Err()
	// if err != nil {
	// 	return nil, err
	// }

	// task, err := DeserializeTask(result)
	// if err != nil {
	// 	return nil, err
	// }

	// var err error
	// var tasksProcessed1 string
	// tasksProcessed := []string{"1"}
	// processedCounter := TasksProcessedCounterPrefix + "1"
	// tasksProcessed1, err = sr.client.Get(processedCounter).Result()
	// if err != nil {
	// 	return nil, err
	// }
	// tasksProcessed[0] = tasksProcessed1

	var tasksScheduled string
	tasksScheduled, err := sr.client.Get(TasksScheduledCounter).Result()
	if err != nil {
		return nil, err
	}

	status := &Status{
		MemoryUsed:     []string{"0"},
		CPUUsed:        []string{"0"},
		DiskUsed:       []string{"0"},
		ProcsRunning:   []string{"0"},
		TasksProcessed: []string{"0"},
		TasksScheduled: tasksScheduled,
	}

	return status, nil
}
