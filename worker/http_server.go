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

func mockResponseHandler(maxValue int) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(common.MockSystemMetric(maxValue)))
	}
}

func listenForHTTPRequests() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/status/healthy", healthyHandler)
	http.HandleFunc("/status/ready", readyHandler)
	http.HandleFunc("/GetMemoryUsed", mockResponseHandler(75))
	http.HandleFunc("/GetCPUUsed", mockResponseHandler(50))
	http.HandleFunc("/GetDiskUsed", mockResponseHandler(25))
	http.HandleFunc("/GetProcsRunning", mockResponseHandler(20))
	http.HandleFunc("/GetGetDiskIO", mockResponseHandler(1000))
	http.HandleFunc("/GetNetworkTraffic", mockResponseHandler(500))
	http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
