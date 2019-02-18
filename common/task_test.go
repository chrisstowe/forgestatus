// +build unit

package common

import (
	"fmt"
	"testing"
)

func TestDeserializeTaskSuccess(t *testing.T) {
	expectedType := GetMemoryUsed
	expectedTime := mockTime
	expectedID := mockID

	expectedTask := Task{
		Type: expectedType,
		Time: expectedTime,
		ID:   expectedID,
	}

	serializedTask := fmt.Sprintf("{\"type\":%q,\"time\":%q,\"id\":%q}",
		expectedType, expectedTime, expectedID)

	actualTask, err := DeserializeTask(serializedTask)
	if err != nil {
		t.Error(err)
	}

	if expectedTask != *actualTask {
		t.Errorf("want: %+v, got: %+v.", expectedTask, *actualTask)
	}
}

func TestDeserializeTaskFailsWithBadData(t *testing.T) {
	serializedTask := fmt.Sprintf("{\"foo\":}")

	_, err := DeserializeTask(serializedTask)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestSerializeTaskSuccess(t *testing.T) {
	expectedType := GetMemoryUsed
	expectedTime := mockTime
	expectedID := mockID

	expectedSerializedTask := fmt.Sprintf("{\"type\":%q,\"time\":%q,\"id\":%q}",
		expectedType, expectedTime, expectedID)

	task := &Task{
		Type: expectedType,
		Time: expectedTime,
		ID:   expectedID,
	}

	actualSerializedTask, err := SerializeTask(task)
	if err != nil {
		t.Error(err)
	}

	if expectedSerializedTask != actualSerializedTask {
		t.Errorf("want: %+v, got: %+v.", expectedSerializedTask, actualSerializedTask)
	}
}
