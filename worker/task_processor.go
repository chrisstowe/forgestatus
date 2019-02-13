package main

import (
	"fmt"
	"time"

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

	fmt.Println(task)

	switch task.Type {
	case common.GetMemoryUsed:
		fmt.Println("getting memory used")
		// m := common.MemoryUsed()
		// fmt.Printf("memory used %f \n", m)
	default:
		fmt.Println("unknown task type")
	}

	// Have the worker backoff
	time.Sleep(time.Second)
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

func startProcessingTasks() {
	defer recoverName()

	for {
		getTasks()
	}
}
