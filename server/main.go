package main

import (
	"fmt"
)

func main() {
	fmt.Println("Started Server")
	scheduleTasks()
	listenForHTTPRequests()
}
