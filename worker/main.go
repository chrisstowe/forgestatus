package main

import (
	"fmt"
	"net/http"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %s", time.Now())
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func main() {
	// common.ExampleNewClient()
	// common.Test()
	// http.HandleFunc("/", timeHandler)
	// http.HandleFunc("/health", healthHandler)
	//http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
