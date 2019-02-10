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

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":80", nil)
}
