package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v! %s", common.Greet("server"), time.Now())
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server is good"))
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
	http.HandleFunc("/", greet)
	http.HandleFunc("/health", health)
	http.HandleFunc("/api/status", status)
	http.ListenAndServe(":80", nil)
}
