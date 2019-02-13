package main

import (
	"fmt"

	"github.com/chrisstowe/forgestatus/common"
)

func getTasks(taskTaker common.TaskTaker) {
	fmt.Println("Getting tasks")

	task, err := taskTaker.TakeNextTask()
	if err != nil {
		fmt.Println("Error taking next task", err)
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
	taskTaker := common.NewTaskTaker(common.EnvConfig.RedisURL)

	err := taskTaker.InitTaskTaker()
	if err != nil {
		fmt.Println("Could not initialize the task taker", err)
		return
	}

	for {
		getTasks(taskTaker)
	}
}
