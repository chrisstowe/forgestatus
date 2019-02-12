package common

import (
	"encoding/json"
)

// TaskType represents the type of work to perform.
type TaskType string

// Enumerated TaskTypes.
const (
	GetUsedMemory        TaskType = "GetUsedMemory"
	GetCpuUsed           TaskType = "GetCpuUsed"
	GetDiskUsed          TaskType = "GetDiskUsed"
	GetProcsRunning      TaskType = "GetProcsRunning"
	GetRequestsProcessed TaskType = "GetRequestsProcessed"
	GetRequestsMade      TaskType = "GetRequestsMade"
)

// Task represents work to be done.
type Task struct {
	Type TaskType `json:"type"`
	Time string   `json:"time"`
}

// DeserializeTask takes a JSON string and converts it to a task.
func DeserializeTask(s string) (Task, error) {
	var task Task
	err := json.Unmarshal([]byte(s), &task)
	if err != nil {
		return Task{}, err
	}

	return task, nil
}

// SerializeTask tasks a task and converts it to a JSON string.
func SerializeTask(task Task) (string, error) {
	b, err := json.Marshal(&task)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
