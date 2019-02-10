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
	//w.Write([]byte("ok"))
	w.WriteHeader(http.StatusInternalServerError)
}

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/health", health)
	http.ListenAndServe(":80", nil)
}
