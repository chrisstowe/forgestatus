package common

import (
	"encoding/json"
)

// PendingQueue is the queue name for pending task work.
const PendingQueue = "pendingQueue"

// ProcessingQueue is the queue name for tasks currently being processed.
// This queue is specific to each worker.
var ProcessingQueue = "processingQueue" + EnvConfig.WorkerID

// ResultQueuePrefix is the prefix for all task result types.
var ResultQueuePrefix = "result"

// TaskType represents the type of work to perform.
type TaskType string

// Enumerated TaskTypes.
const (
	GetMemoryUsed        TaskType = "GetMemoryUsed"
	GetCPUUsed           TaskType = "GetCPUUsed"
	GetDiskUsed          TaskType = "GetDiskUsed"
	GetProcsRunning      TaskType = "GetProcsRunning"
	GetRequestsProcessed TaskType = "GetRequestsProcessed"
	GetRequestsMade      TaskType = "GetRequestsMade"
)

// Task represents work to be done.
type Task struct {
	Type TaskType `json:"type"`
	Time string   `json:"time"`
	ID   string   `json:"id"`
	Data string   `json:"data"`
}

// DeserializeTask takes a JSON string and converts it to a task.
func DeserializeTask(s string) (*Task, error) {
	task := &Task{}
	err := json.Unmarshal([]byte(s), task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

// SerializeTask tasks a task and converts it to a JSON string.
func SerializeTask(task *Task) (string, error) {
	b, err := json.Marshal(task)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
