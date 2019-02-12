package main

import "github.com/chrisstowe/forgestatus/common"

func main() {
	go startSchedulingTasks()
	listenForHttpRequests(common.EnvConfig.Port)
}
