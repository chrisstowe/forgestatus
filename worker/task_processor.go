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
		r := common.NewResult(task.Type, task.ID, []string{"91"})
		taskTaker.SetTaskResult(r)
	case common.GetCPUUsed:
		r := common.NewResult(task.Type, task.ID, []string{"44"})
		taskTaker.SetTaskResult(r)
	case common.GetDiskUsed:
		r := common.NewResult(task.Type, task.ID, []string{"33"})
		taskTaker.SetTaskResult(r)
	case common.GetProcsRunning:
		r := common.NewResult(task.Type, task.ID, []string{"3"})
		taskTaker.SetTaskResult(r)
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
