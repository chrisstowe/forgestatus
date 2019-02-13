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

	taskScheduler.ScheduleTask(task)
}

func recoverName() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from ", r)
	}
}

func startSchedulingTasks() {
	defer recoverName()

	ticker := time.NewTicker(200 * time.Millisecond)
	for range ticker.C {
		scheduledTasks()
	}
}
