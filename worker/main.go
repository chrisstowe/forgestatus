package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %s", time.Now())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func getTasks() {
	taskTaker := common.NewTaskTaker(common.EnvConfig.RedisURL)

	task, err := taskTaker.TakeNextTask()
	if err != nil {
		return
	}

	fmt.Println(task)

	switch task.Type {
	case common.GetMemoryUsed:
		fmt.Println("getting memory used")
	default:
		fmt.Println("unkown task type")
	}
}

func main() {
	// for {
	// 	getTasks()
	// }
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
