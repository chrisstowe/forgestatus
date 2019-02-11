package common

import "os"

// Config holds required and optional config parameters.
type Config struct {
	Port,
	WorkerID string
}

// EnvConfig holds the current environment configs.
var EnvConfig = Config{
	Port:     os.Getenv("PORT"),
	WorkerID: os.Getenv("WORKER_ID"),
}
