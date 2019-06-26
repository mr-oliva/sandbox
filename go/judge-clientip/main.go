package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.Header.Get("X-Forwarded-For")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"ip": "%s"}`, clientIP)
	})

    http.ListenAndServe(":8080", nil)
}
