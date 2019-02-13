package main

import (
	"fmt"
	"time"

	"github.com/rs/xid"

	"github.com/chrisstowe/forgestatus/common"
)

var taskScheduler = common.NewTaskScheduler(common.EnvConfig.RedisURL)

func scheduledTasks() {
	fmt.Println("Scheduling tasks")

	task := &common.Task{
		Type: common.GetMemoryUsed,
		Time: time.Now().Format(time.RFC3339Nano),
		ID:   xid.New().String(),
	}

	err := taskScheduler.ScheduleTask(task)
	if err != nil {
		fmt.Println("Error scheduling tasks", err)
		return
	}
}

func startSchedulingTasks() {
	ticker := time.NewTicker(time.Second)
	for range ticker.C {
		scheduledTasks()
	}
}
