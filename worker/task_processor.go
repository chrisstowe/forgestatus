package main

import (
	"fmt"

	"github.com/chrisstowe/forgestatus/common"
)

var taskTaker = common.NewTaskTaker(common.EnvConfig.RedisURL)

func getTasks() {
	fmt.Println("Getting tasks")

	task, err := taskTaker.TakeNextTask()
	if err != nil {
		fmt.Println("Error taking next task", err)
		return
	}

	fmt.Printf("Got task: %+v\n", task)

	switch task.Type {
	case common.GetMemoryUsed:
		fmt.Println("mem")
	case common.GetCPUUsed:
		fmt.Println("cpu")
	case common.GetDiskUsed:
		fmt.Println("disk")
	case common.GetProcsRunning:
		fmt.Println("procs")
	default:
		fmt.Println("Unknown task type")
	}
}

func processTasks() {
	go func() {
		for {
			getTasks()
		}
	}()
}
