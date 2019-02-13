package common

import (
	"encoding/json"
	"time"

	"github.com/rs/xid"
)

// Task represents work to be done.
type Task struct {
	Type TaskType `json:"type"`
	Time string   `json:"time"`
	ID   string   `json:"id"`
}

// NewTask creates a new Task.
func NewTask(t TaskType) *Task {
	return &Task{
		Type: t,
		Time: time.Now().Format(time.RFC3339Nano),
		ID:   xid.New().String(),
	}
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

// SerializeTask takes a task and converts it to a JSON string.
func SerializeTask(t *Task) (string, error) {
	b, err := json.Marshal(t)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
