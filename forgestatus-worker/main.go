package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

var workerID = os.Getenv("WORKER_ID")

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v! %s", common.Greet("worker"), time.Now())
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "worker %v is good", workerID)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/health", health)
	http.HandleFunc("/status", status)
	http.ListenAndServe(":80", nil)
}
