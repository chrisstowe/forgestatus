package main

import (
	"fmt"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

var taskScheduler = common.NewTaskScheduler(common.EnvConfig.RedisURL)

func scheduleTask(t common.TaskType) {
	fmt.Println("Scheduling task: ", t)

	task := common.NewTask(t, "")

	err := taskScheduler.ScheduleTask(task)
	if err != nil {
		fmt.Println("Error scheduling tasks", err)
	}
}

func scheduleTasks() {
	common.Schedule(func() { scheduleTask(common.GetMemoryUsed) }, time.Second)
	common.Schedule(func() { scheduleTask(common.GetCPUUsed) }, time.Second)
	common.Schedule(func() { scheduleTask(common.GetDiskUsed) }, time.Second)
	common.Schedule(func() { scheduleTask(common.GetProcsRunning) }, time.Second)
}
