package common

import "github.com/go-redis/redis"

// TaskScheduler schedules tasks to be worked on.
type TaskScheduler interface {
	ScheduleTask(*Task) error
}

type taskScheduler struct {
	client *redis.Client
}

// NewTaskScheduler creates a new TaskScheduler.
func NewTaskScheduler(redisURL string) TaskScheduler {
	c := redis.NewClient(&redis.Options{Addr: redisURL})
	return &taskScheduler{client: c}
}

func (ts *taskScheduler) ScheduleTask(task *Task) error {
	s, err := SerializeTask(task)
	if err != nil {
		return err
	}

	err = ts.client.LPush(PendingQueue, s).Err()
	if err != nil {
		return err
	}

	// Prevent too many tasks from building up.
	// In a real system, the number of workers should probably be scaled up.
	// This is an O(1) operation (since the worst case is always removing 1).
	err = ts.client.LTrim(PendingQueue, 0, EnvConfig.MaxTaskQueueSize-1).Err()
	if err != nil {
		return err
	}

	// Increment the amount of tasks that have been scheduled.
	err = ts.client.Incr(TasksScheduledCounter).Err()
	if err != nil {
		return err
	}

	return nil
}
