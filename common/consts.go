package common

// TaskType represents the type of work to perform.
type TaskType string

// Enumerated TaskTypes.
const (
	GetMemoryUsed     TaskType = "GetMemoryUsed"
	GetCPUUsed        TaskType = "GetCPUUsed"
	GetDiskUsed       TaskType = "GetDiskUsed"
	GetProcsRunning   TaskType = "GetProcsRunning"
	GetDiskIO         TaskType = "GetDiskIO"
	GetNetworkTraffic TaskType = "GetNetworkTraffic"
)

const (
	// PendingQueue is the queue name for pending task work.
	PendingQueue = "pendingQueue"

	// ProcessingQueuePrefix is the queue prefix for tasks currently being processed.
	// This queue is specific to each worker.
	ProcessingQueuePrefix = "processingQueue"

	// ResultQueuePrefix is the prefix for all task result types.
	ResultQueuePrefix = "resultQueue"

	// TasksScheduledCounter is the counter name for the number of scheduled tasks.
	TasksScheduledCounter = "tasksScheduledCounter"

	// TasksProcessedCounterPrefix is the counter name for the number of completed tasks.
	// This counter is specific to each worker.
	TasksProcessedCounterPrefix = "tasksProcessedCount"
)
