package common

import (
	"fmt"
	"testing"
)

const mockTime = "2019-02-11T18:27:02.566923-08:00"
const mockID = "9m4e2mr0ui3e8a215n4g"

func TestDeserializeTaskSuccess(t *testing.T) {
	expectedType := GetMemoryUsed
	expectedTime := mockTime
	expectedID := mockID
	expectedData := "foo"

	expectedTask := Task{
		Type: expectedType,
		Time: expectedTime,
		ID:   expectedID,
		Data: expectedData,
	}

	serializedTask := fmt.Sprintf("{\"type\":%q,\"time\":%q,\"id\":%q,\"data\":%q}",
		expectedType, expectedTime, expectedID, expectedData)

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
	expectedID := mockID
	expectedData := "foo"

	expectedSerializedTask := fmt.Sprintf("{\"type\":%q,\"time\":%q,\"id\":%q,\"data\":%q}",
		expectedType, expectedTime, expectedID, expectedData)

	task := Task{
		Type: expectedType,
		Time: expectedTime,
		ID:   expectedID,
		Data: expectedData,
	}

	actualSerializedTask, err := SerializeTask(task)
	if err != nil {
		t.Error(err)
	}

	if expectedSerializedTask != actualSerializedTask {
		t.Errorf("want: %+v, got: %+v.", expectedSerializedTask, actualSerializedTask)
	}
}
