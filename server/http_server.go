package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

var statusReader = common.NewStatusReader(common.EnvConfig.RedisURL)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Time: %s", time.Now())
}

func healthyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("healthy"))
}

func readyHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ready"))
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	status, err := statusReader.GetStatus()
	if err != nil {
		http.Error(w, "Problem reading status",
			http.StatusInternalServerError)
	}

	s, err := common.SerializeStatus(status)
	if err != nil {
		http.Error(w, "Problem serializing status",
			http.StatusInternalServerError)
	}

	w.Write([]byte(s))
}

func listenForHTTPRequests() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/status/healthy", healthyHandler)
	http.HandleFunc("/status/ready", readyHandler)
	http.HandleFunc("/api/status", statusHandler)
	http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
