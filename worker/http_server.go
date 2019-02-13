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

func getMemoryUsedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(common.MockSystemMetric(80)))
}

func getCPUUsedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(common.MockSystemMetric(50)))
}

func getDiskUsedHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(common.MockSystemMetric(20)))
}

func getProcsRunningHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(common.MockSystemMetric(100)))
}

func listenForHTTPRequests() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/GetMemoryUsed", getMemoryUsedHandler)
	http.HandleFunc("/GetCPUUsed", getCPUUsedHandler)
	http.HandleFunc("/GetDiskUsed", getDiskUsedHandler)
	http.HandleFunc("/GetProcsRunning", getProcsRunningHandler)
	http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
