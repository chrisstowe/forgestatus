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
		t := common.NewTask(task.Type, "32.0")
		taskTaker.SetTaskResult(t)
	case common.GetCPUUsed:
		t := common.NewTask(task.Type, "13.0")
		taskTaker.SetTaskResult(t)
	case common.GetDiskUsed:
		t := common.NewTask(task.Type, "19.0")
		taskTaker.SetTaskResult(t)
	case common.GetProcsRunning:
		t := common.NewTask(task.Type, "3")
		taskTaker.SetTaskResult(t)
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
