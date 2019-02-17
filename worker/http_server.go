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

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ready"))
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
	http.HandleFunc("/status/healthy", healthyHandler)
	http.HandleFunc("/status/ready", readyHandler)
	http.HandleFunc("/GetMemoryUsed", getMemoryUsedHandler)
	http.HandleFunc("/GetCPUUsed", getCPUUsedHandler)
	http.HandleFunc("/GetDiskUsed", getDiskUsedHandler)
	http.HandleFunc("/GetProcsRunning", getProcsRunningHandler)
	http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
