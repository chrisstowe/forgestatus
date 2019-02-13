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

func listenForHTTPRequests(port string) {
	http.HandleFunc("/", timeHandler)
	http.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":"+port, nil)
}
