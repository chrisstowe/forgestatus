package main

import "github.com/chrisstowe/forgestatus/common"

func main() {
	go startProcessingTasks()
	listenForHttpRequests(common.EnvConfig.Port)
}
