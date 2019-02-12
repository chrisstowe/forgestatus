package main

import (
	"fmt"

	"github.com/chrisstowe/forgestatus/common"
)

func getTasks() {
	taskTaker := common.NewTaskTaker(common.EnvConfig.RedisURL)

	task, err := taskTaker.TakeNextTask()
	if err != nil {
		return
	}

	fmt.Println(task)

	switch task.Type {
	case common.GetMemoryUsed:
		fmt.Println("getting memory used")
		// m := common.MemoryUsed()
		// fmt.Printf("memory used %f \n", m)
	default:
		fmt.Println("unknown task type")
	}
}

func startProcessingTasks() {
	for {
		getTasks()
	}
}
