package common

import (
	"github.com/go-redis/redis"
)

// TaskTaker takes tasks and sets the results.
type TaskTaker interface {
	TakeNextTask() (*Task, error)
	SetTaskResult(*Task) error
}

type taskTaker struct {
	client *redis.Client
}

// NewTaskTaker creates a new TaskTaker
func NewTaskTaker(redisURL string) TaskTaker {
	c := redis.NewClient(&redis.Options{Addr: redisURL})
	return &taskTaker{client: c}
}

func (tt *taskTaker) TakeNextTask() (*Task, error) {
	// This queue is specific to each worker.
	processingQueue := ProcessingTaskQueue + EnvConfig.WorkerID

	// Move a task from the pending queue into a processing one.
	// This is an atomic operation, so no data is lost.
	result, err := tt.client.BRPopLPush(PendingTaskQueue, processingQueue, 0).Result()
	if err != nil {
		return nil, err
	}

	// One task should be worked on at a time.
	// This is an arbitrary limitation that can be improved later.
	// There is also no process for pushing stale tasks back into the pending queue.
	// This is an O(1) operation (since the worst case is always removing 1).
	err = tt.client.LTrim(processingQueue, 0, 0).Err()
	if err != nil {
		return nil, err
	}

	task, err := DeserializeTask(result)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (tt *taskTaker) SetTaskResult(task *Task) error {
	return nil
}
