package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeserializeResultSuccess(t *testing.T) {
	expectedType := GetMemoryUsed
	expectedTime := mockTime
	expectedID := mockID
	expectedData := "32"

	expectedResult := Result{
		Type: expectedType,
		Time: expectedTime,
		ID:   expectedID,
		Data: []string{expectedData},
	}

	serializedResult := fmt.Sprintf("{\"type\":%q,\"time\":%q,\"id\":%q,\"data\":[%q]}",
		expectedType, expectedTime, expectedID, expectedData)

	actualResult, err := DeserializeResult(serializedResult)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expectedResult, *actualResult) {
		t.Errorf("want: %+v, got: %+v.", expectedResult, *actualResult)
	}
}

func TestDeserializeResultFailsWithBadData(t *testing.T) {
	serializedResult := fmt.Sprintf("{\"foo\":}")

	_, err := DeserializeResult(serializedResult)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestSerializeResultSuccess(t *testing.T) {
	expectedType := GetMemoryUsed
	expectedTime := mockTime
	expectedID := mockID
	expectedData := "32"

	expectedSerializedResult := fmt.Sprintf("{\"type\":%q,\"time\":%q,\"id\":%q,\"data\":[%q]}",
		expectedType, expectedTime, expectedID, expectedData)

	result := &Result{
		Type: expectedType,
		Time: expectedTime,
		ID:   expectedID,
		Data: []string{expectedData},
	}

	actualSerializedResult, err := SerializeResult(result)
	if err != nil {
		t.Error(err)
	}

	if expectedSerializedResult != actualSerializedResult {
		t.Errorf("want: %+v, got: %+v.", expectedSerializedResult, actualSerializedResult)
	}
}
