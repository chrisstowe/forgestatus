// +build integration

package common

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_ScheduleTask_SerializesCorrectly(t *testing.T) {
	redisTester, err := NewRedisTester()
	if err != nil {
		t.Fatal(err)
	}
	defer redisTester.Close()

	taskScheduler := NewTaskScheduler(TestRedisURL)

	for _, taskType := range taskTypes {
		fmt.Println("testing task: ", taskType)

		expectedTask := NewTask(taskType)

		err = taskScheduler.ScheduleTask(expectedTask)
		if err != nil {
			t.Fatal(err)
		}

		serializedTask, err := redisTester.Client.RPop(PendingQueue).Result()
		if err != nil {
			t.Fatal(err)
		}

		actualTask, err := DeserializeTask(serializedTask)
		if err != nil {
			t.Fatal(err)
		}

		// ASSERT
		if *expectedTask != *actualTask {
			t.Errorf("want: %+v, got: %+v.", *expectedTask, *actualTask)
		}
	}
}

func Test_ScheduleTask_RespectsMaxTaskQueueSize(t *testing.T) {
	redisTester, err := NewRedisTester()
	if err != nil {
		t.Fatal(err)
	}
	defer redisTester.Close()

	expectedTaskQueueSize := 20

	// Restrict the task scheduler to this many pending tasks.
	EnvConfig.MaxTaskQueueSize = int64(expectedTaskQueueSize)

	// Schedule more than the max amount of tasks.
	taskCountToSchedule := expectedTaskQueueSize + 10

	taskScheduler := NewTaskScheduler(TestRedisURL)

	for i := 0; i < taskCountToSchedule; i++ {
		expectedTask := NewTask(GetHealthy)

		err = taskScheduler.ScheduleTask(expectedTask)
		if err != nil {
			t.Fatal(err)
		}
	}

	tasks, err := redisTester.Client.LRange(PendingQueue, 0, int64(taskCountToSchedule-1)).Result()
	if err != nil {
		t.Fatal(err)
	}

	actualTaskQueueSize := len(tasks)

	// ASSERT
	if expectedTaskQueueSize != actualTaskQueueSize {
		t.Errorf("want: %+v, got: %+v.", expectedTaskQueueSize, actualTaskQueueSize)
	}
}

func Test_ScheduleTask_ScheduledCounterIsIncremented(t *testing.T) {
	redisTester, err := NewRedisTester()
	if err != nil {
		t.Fatal(err)
	}
	defer redisTester.Close()

	expectedTaskCount := 10

	taskScheduler := NewTaskScheduler(TestRedisURL)

	for i := 0; i < expectedTaskCount; i++ {
		expectedTask := NewTask(GetHealthy)

		err = taskScheduler.ScheduleTask(expectedTask)
		if err != nil {
			t.Fatal(err)
		}
	}

	taskCount, err := redisTester.Client.Get(TasksScheduledCounter).Result()
	if err != nil {
		t.Fatal(err)
	}

	actualTaskCount, err := strconv.Atoi(taskCount)
	if err != nil {
		t.Fatal(err)
	}

	// ASSERT
	if expectedTaskCount != actualTaskCount {
		t.Errorf("want: %+v, got: %+v.", expectedTaskCount, actualTaskCount)
	}
}
