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

func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

// func getWorkerStatus(w http.ResponseWriter, r *http.Request) {
// 	resp, err := http.Get("http://forgestatus-worker-1-service-dev/getStatus")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Could not connect to the worker!"))
// 		return
// 	}

// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Write(body)
// }

func main() {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/api/status", statusHandler)
	http.ListenAndServe(":"+common.EnvConfig.Port, nil)
}
