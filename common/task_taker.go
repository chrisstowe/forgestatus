package common

import "github.com/go-redis/redis"

// TaskTaker takes tasks and sets the results.
type TaskTaker interface {
	TakeNextTask() (*Task, error)
	SetTaskResult(*Result) error
}

type taskTaker struct {
	client *redis.Client
}

var processingQueue = ProcessingQueuePrefix + EnvConfig.WorkerID
var tasksProcessedCounter = TasksProcessedCounterPrefix + EnvConfig.WorkerID

// NewTaskTaker creates a new TaskTaker.
func NewTaskTaker(redisURL string) TaskTaker {
	c := redis.NewClient(&redis.Options{Addr: redisURL})
	return &taskTaker{client: c}
}

func (tt *taskTaker) TakeNextTask() (*Task, error) {
	// Move a task from the pending queue into a processing one.
	// This is an atomic operation, so no data is lost.
	result, err := tt.client.BRPopLPush(PendingQueue, processingQueue, 0).Result()
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

	return task, nil
}

func (tt *taskTaker) SetTaskResult(result *Result) error {
	s, err := SerializeResult(result)
	if err != nil {
		return err
	}

	// Push to the queue related to the task type.
	resultQueue := ResultQueuePrefix + string(result.Type)
	err = tt.client.LPush(resultQueue, s).Err()
	if err != nil {
		return err
	}

	// Prevent too many results from building up.
	// This is an O(1) operation (since the worst case is always removing 1).
	err = tt.client.LTrim(resultQueue, 0, EnvConfig.MaxResultQueueSize-1).Err()
	if err != nil {
		return err
	}

	// One task should be worked on at a time.
	// This is an arbitrary limitation that can be improved later.
	// Remove the last task in the processing queue.
	// LREM could also be used here, but the original task would need passed in.
	err = tt.client.RPop(processingQueue).Err()
	if err != nil {
		return err
	}

	// Increment the amount of tasks that have been proccessed.
	err = tt.client.Incr(tasksProcessedCounter).Err()
	if err != nil {
		return err
	}

	return nil
}
