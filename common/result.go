package common

import (
	"encoding/json"
	"time"
)

// Result represents the results from a task.
type Result struct {
	Type TaskType `json:"type"`
	Time string   `json:"time"`
	ID   string   `json:"id"`
	Data []string `json:"data"`
}

// NewResult creates a new Result.
func NewResult(t TaskType, id string, data []string) *Result {
	return &Result{
		Type: t,
		Time: time.Now().Format(time.RFC3339Nano),
		ID:   id,
		Data: data,
	}
}

// DeserializeResult takes a JSON string and converts it to a result.
func DeserializeResult(s string) (*Result, error) {
	result := &Result{}
	err := json.Unmarshal([]byte(s), result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// SerializeResult takes a result and converts it to a JSON string.
func SerializeResult(r *Result) (string, error) {
	b, err := json.Marshal(r)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
