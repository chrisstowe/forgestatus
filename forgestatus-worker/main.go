package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/chrisstowe/forgestatus/common"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %v! %s", common.Greet("worker"), time.Now())
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("worker is good"))
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/health", health)
	http.HandleFunc("/worker1/status", status)
	http.HandleFunc("/worker2/status", status)
	http.HandleFunc("/worker3/status", status)
	http.ListenAndServe(":80", nil)
}
