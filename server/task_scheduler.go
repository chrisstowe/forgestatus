package main

import (
	"time"

	"github.com/rs/xid"

	"github.com/chrisstowe/forgestatus/common"
)

func scheduledTasks() {
	taskScheduler := common.NewTaskScheduler(common.EnvConfig.RedisURL)

	task := common.Task{
		Type: common.GetMemoryUsed,
		Time: time.Now().Format(time.RFC3339Nano),
		ID:   xid.New().String(),
	}

	taskScheduler.ScheduleTask(task)
}

func startSchedulingTasks() {
	ticker := time.NewTicker(200 * time.Millisecond)
	for range ticker.C {
		scheduledTasks()
	}
}
