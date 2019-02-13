package common

import "encoding/json"

// Status represents an aggregation of the cluster status.
type Status struct {
	MemoryUsed     []string `json:"memoryUsed"`
	CPUUsed        []string `json:"cpuUsed"`
	DiskUsed       []string `json:"diskUsed"`
	ProcsRunning   []string `json:"procsRunning"`
	TasksProcessed []string `json:"tasksProcessed"`
	TasksScheduled string   `json:"tasksScheduled"`
}

// DeserializeStatus takes a JSON string and converts it to a status.
func DeserializeStatus(s string) (*Status, error) {
	status := &Status{}
	err := json.Unmarshal([]byte(s), status)
	if err != nil {
		return nil, err
	}

	return status, nil
}

// SerializeStatus takes a status and converts it to a JSON string.
func SerializeStatus(status *Status) (string, error) {
	b, err := json.Marshal(status)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
