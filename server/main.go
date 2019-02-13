package main

import (
	"fmt"

	"github.com/chrisstowe/forgestatus/common"
)

func main() {
	fmt.Println("Started Server")
	go startSchedulingTasks()
	listenForHTTPRequests(common.EnvConfig.Port)
}
