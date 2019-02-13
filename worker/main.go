package main

import (
	"fmt"

	"github.com/chrisstowe/forgestatus/common"
)

func main() {
	fmt.Println("Started Worker ", common.EnvConfig.WorkerID)
	go startProcessingTasks()
	listenForHTTPRequests()
}
