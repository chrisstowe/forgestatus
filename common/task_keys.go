package common

var (
	// PendingQueue is the queue name for pending task work.
	PendingQueue = "pendingQueue"

	// ProcessingQueue is the queue name for tasks currently being processed.
	// This queue is specific to each worker.
	ProcessingQueue = "processingQueue" + EnvConfig.WorkerID

	// ResultQueuePrefix is the prefix for all task result types.
	ResultQueuePrefix = "result"

	// TasksScheduledCounter is the counter name for the number of scheduled tasks.
	TasksScheduledCounter = "tasksScheduledCounter"

	// TasksProcessedCounter is the counter name for the number of completed tasks.
	// This counter is specific to each worker.
	TasksProcessedCounter = "tasksProcessedCount" + EnvConfig.WorkerID
)
