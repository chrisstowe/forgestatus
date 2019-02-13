package common

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeserializeStatusSuccess(t *testing.T) {
	expectedStatus := Status{
		MemoryUsed:     []string{"21"},
		CPUUsed:        []string{"23"},
		DiskUsed:       []string{"11"},
		ProcsRunning:   []string{"15"},
		TasksProcessed: []string{"1000"},
		TasksScheduled: "3000",
	}

	serializedStatus := "{\"memoryUsed\":[\"21\"],\"cpuUsed\":[\"23\"],\"diskUsed\":[\"11\"],\"procsRunning\":[\"15\"],\"tasksProcessed\":[\"1000\"],\"tasksScheduled\":\"3000\"}"

	actualStatus, err := DeserializeStatus(serializedStatus)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expectedStatus, *actualStatus) {
		t.Errorf("want: %+v, got: %+v.", expectedStatus, *actualStatus)
	}
}

func TestDeserializeStatusFailsWithBadData(t *testing.T) {
	serializedStatus := fmt.Sprintf("{\"foo\":}")

	_, err := DeserializeStatus(serializedStatus)
	if err == nil {
		t.Error("Expected an error")
	}
}

func TestSerializeStatusSuccess(t *testing.T) {
	expectedSerializedStatus := "{\"memoryUsed\":[\"21\"],\"cpuUsed\":[\"23\"],\"diskUsed\":[\"11\"],\"procsRunning\":[\"15\"],\"tasksProcessed\":[\"1000\"],\"tasksScheduled\":\"3000\"}"

	status := &Status{
		MemoryUsed:     []string{"21"},
		CPUUsed:        []string{"23"},
		DiskUsed:       []string{"11"},
		ProcsRunning:   []string{"15"},
		TasksProcessed: []string{"1000"},
		TasksScheduled: "3000",
	}

	actualSerializedStatus, err := SerializeStatus(status)
	if err != nil {
		t.Error(err)
	}

	if expectedSerializedStatus != actualSerializedStatus {
		t.Errorf("want: %+v, got: %+v.", expectedSerializedStatus, actualSerializedStatus)
	}
}
