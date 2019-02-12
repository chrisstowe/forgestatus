package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s", time.Now())
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "worker %v is good", common.EnvConfig.WorkerID)
}

func main() {
	// common.ExampleNewClient()
	// common.Test()
	// http.HandleFunc("/", greet)
	// http.HandleFunc("/health", health)
	// http.HandleFunc("/status", status)
	//http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
