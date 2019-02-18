// +build integration

package common

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_TakeNextTask_DeserializesCorrectly(t *testing.T) {
	redisTester, err := NewRedisTester()
	if err != nil {
		t.Fatal(err)
	}
	defer redisTester.Close()

	taskTaker := NewTaskTaker(TestRedisURL)

	for _, taskType := range taskTypes {
		fmt.Println("testing task: ", taskType)

		expectedTask := NewTask(taskType)

		serializedTask, err := SerializeTask(expectedTask)
		if err != nil {
			t.Fatal(err)
		}

		err = redisTester.Client.LPush(PendingQueue, serializedTask).Err()
		if err != nil {
			t.Fatal(err)
		}

		actualTask, err := taskTaker.TakeNextTask()
		if err != nil {
			t.Fatal(err)
		}

		// ASSERT
		if *expectedTask != *actualTask {
			t.Errorf("want: %+v, got: %+v.", *expectedTask, *actualTask)
		}
	}
}

func Test_TakeNextTask_MovesTaskFromPendingToProcessingQueue(t *testing.T) {
	redisTester, err := NewRedisTester()
	if err != nil {
		t.Fatal(err)
	}
	defer redisTester.Close()

	processingQueue := ProcessingQueuePrefix + strconv.Itoa(EnvConfig.WorkerID)
	if err != nil {
		t.Fatal(err)
	}

	taskTaker := NewTaskTaker(TestRedisURL)

	for _, taskType := range taskTypes {
		fmt.Println("testing task: ", taskType)

		expectedTask := NewTask(taskType)

		serializedExpectedTask, err := SerializeTask(expectedTask)
		if err != nil {
			t.Fatal(err)
		}

		err = redisTester.Client.LPush(PendingQueue, serializedExpectedTask).Err()
		if err != nil {
			t.Fatal(err)
		}

		_, err = taskTaker.TakeNextTask()
		if err != nil {
			t.Fatal(err)
		}

		serializedActualTask, err := redisTester.Client.RPop(processingQueue).Result()
		if err != nil {
			t.Fatal(err)
		}

		actualTask, err := DeserializeTask(serializedActualTask)
		if err != nil {
			t.Fatal(err)
		}

		// ASSERT
		if *expectedTask != *actualTask {
			t.Errorf("want: %+v, got: %+v.", *expectedTask, *actualTask)
		}
	}
}

////////

func Test_SetTaskResult_SerializesCorrectly(t *testing.T) {
	redisTester, err := NewRedisTester()
	if err != nil {
		t.Fatal(err)
	}
	defer redisTester.Close()

	processingQueue := ProcessingQueuePrefix + strconv.Itoa(EnvConfig.WorkerID)
	if err != nil {
		t.Fatal(err)
	}

	taskTaker := NewTaskTaker(TestRedisURL)

	for _, taskType := range taskTypes {
		fmt.Println("testing result for task: ", taskType)

		// A task must be in the processing queue before a result can be set.
		task := NewTask(taskType)

		serializedTask, err := SerializeTask(task)
		if err != nil {
			t.Fatal(err)
		}

		err = redisTester.Client.LPush(processingQueue, serializedTask).Err()
		if err != nil {
			t.Fatal(err)
		}

		// Add a mock result.
		expectedResult := NewResult(taskType, "1", []string{"1", "2", "3"})

		err = taskTaker.SetTaskResult(expectedResult)
		if err != nil {
			t.Fatal(err)
		}

		resultQueue := ResultQueuePrefix + string(taskType)

		serializedResult, err := redisTester.Client.RPop(resultQueue).Result()
		if err != nil {
			t.Fatal(err)
		}

		actualResult, err := DeserializeResult(serializedResult)
		if err != nil {
			t.Fatal(err)
		}

		// ASSERT
		if !reflect.DeepEqual(*expectedResult, *actualResult) {
			t.Errorf("want: %+v, got: %+v.", *expectedResult, *actualResult)
		}
	}
}

func Test_SetTaskResult_RespectsMaxResultQueueSize(t *testing.T) {
	redisTester, err := NewRedisTester()
	if err != nil {
		t.Fatal(err)
	}
	defer redisTester.Close()

	taskType := GetHealthy

	resultQueue := ResultQueuePrefix + string(taskType)

	processingQueue := ProcessingQueuePrefix + strconv.Itoa(EnvConfig.WorkerID)
	if err != nil {
		t.Fatal(err)
	}

	expectedResultQueueSize := 20

	// Restrict the task taker to this many results.
	EnvConfig.MaxResultQueueSize = int64(expectedResultQueueSize)

	// Set more than the max amount of results.
	resultCountToSet := expectedResultQueueSize + 10

	taskTaker := NewTaskTaker(TestRedisURL)

	for i := 0; i < resultCountToSet; i++ {
		// A task must be in the processing queue before a result can be set.
		task := NewTask(taskType)

		serializedTask, err := SerializeTask(task)
		if err != nil {
			t.Fatal(err)
		}

		err = redisTester.Client.LPush(processingQueue, serializedTask).Err()
		if err != nil {
			t.Fatal(err)
		}

		// Add a mock result.
		expectedTask := NewResult(taskType, "1", []string{"1", "2", "3"})

		err = taskTaker.SetTaskResult(expectedTask)
		if err != nil {
			t.Fatal(err)
		}
	}

	results, err := redisTester.Client.LRange(resultQueue, 0, int64(resultCountToSet-1)).Result()
	if err != nil {
		t.Fatal(err)
	}

	actualResultQueueSize := len(results)

	// ASSERT
	if expectedResultQueueSize != actualResultQueueSize {
		t.Errorf("want: %+v, got: %+v.", expectedResultQueueSize, actualResultQueueSize)
	}
}
