package main

import (
	"net/http"

	"github.com/bookun/sandbox/go/judge-clientip/function"
)

func main() {
	http.HandleFunc("/", function.GetIP)

	http.ListenAndServe(":8080", nil)
}
