package common

import (
	"fmt"
	"testing"
)

const mockTime = "2019-02-11T18:27:02.566923-08:00"

func TestDeserializeTaskSuccess(t *testing.T) {
	expectedType := GetMemoryUsed
	expectedTime := mockTime

	expectedTask := Task{
		Type: expectedType,
		Time: expectedTime,
	}

	serializedTask := fmt.Sprintf("{\"type\":\"%s\",\"time\":\"%s\"}",
		expectedType, expectedTime)

	actualTask, err := DeserializeTask(serializedTask)
	if err != nil {
		t.Error(err)
	}

	if expectedTask != actualTask {
		t.Errorf("want: %+v, got: %+v.", expectedTask, actualTask)
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

	expectedSerializedTask := fmt.Sprintf("{\"type\":\"%s\",\"time\":\"%s\"}",
		expectedType, expectedTime)

	task := Task{
		Type: expectedType,
		Time: expectedTime,
	}

	actualSerializedTask, err := SerializeTask(task)
	if err != nil {
		t.Error(err)
	}

	if expectedSerializedTask != actualSerializedTask {
		t.Errorf("want: %+v, got: %+v.", expectedSerializedTask, actualSerializedTask)
	}
}
