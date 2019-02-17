package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "time: %s", time.Now())
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ready"))
}

func mockHealthyReadyHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(common.MockHealthyReady()))
	}
}

func mockSystemValueHandler(maxValue int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(common.MockSystemValue(maxValue)))
	}
}

func listenForHTTPRequests() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/status/healthy", healthyHandler)
	http.HandleFunc("/status/ready", readyHandler)

	// Simulated system metrics.
	http.HandleFunc("/GetHealthy", mockHealthyReadyHandler())
	http.HandleFunc("/GetReady", mockHealthyReadyHandler())
	http.HandleFunc("/GetMemoryUsed", mockSystemValueHandler(75))
	http.HandleFunc("/GetCPUUsed", mockSystemValueHandler(50))
	http.HandleFunc("/GetDiskUsed", mockSystemValueHandler(25))
	http.HandleFunc("/GetProcsRunning", mockSystemValueHandler(20))
	http.HandleFunc("/GetGetDiskIO", mockSystemValueHandler(1000))
	http.HandleFunc("/GetNetworkTraffic", mockSystemValueHandler(500))

	http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
