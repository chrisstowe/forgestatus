package common

import "encoding/json"

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

// SerializeTask takes a task and converts it to a JSON string.
func SerializeTask(task *Task) (string, error) {
	b, err := json.Marshal(task)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
