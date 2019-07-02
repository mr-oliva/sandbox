package main

import (
	"net/http"

	"github.com/bookun/sandbox/go/judge-clientip/ip"
)

func main() {
	http.HandleFunc("/", ip.GetIP)

	http.ListenAndServe(":8080", nil)
}
