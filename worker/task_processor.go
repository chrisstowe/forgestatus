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

	status := getStatusForAllWorkers(task.Type)

	fmt.Printf("Got status for %s: %+v\n", task.Type, task)

	result := common.NewResult(task.Type, task.ID, status)
	taskTaker.SetTaskResult(result)
}

func processTasks() {
	go func() {
		for {
			getTasks()
		}
	}()
}
