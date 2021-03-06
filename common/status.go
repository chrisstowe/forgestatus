package common

import "encoding/json"

// Status represents an aggregation of the cluster status.
type Status struct {
	TasksScheduled string   `json:"tasksScheduled"`
	TasksProcessed []string `json:"tasksProcessed"`
	Healthy        []Result `json:"healthy"`
	Ready          []Result `json:"ready"`
	MemoryUsed     []Result `json:"memoryUsed"`
	CPUUsed        []Result `json:"cpuUsed"`
	DiskUsed       []Result `json:"diskUsed"`
	ProcsRunning   []Result `json:"procsRunning"`
	DiskIO         []Result `json:"diskIO"`
	NetworkTraffic []Result `json:"networkTraffic"`
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
func SerializeStatus(s *Status) (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
